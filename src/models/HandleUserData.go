package models

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"

	PasswordHandler "../password"
)

const userImgDir = "./img/user_img/"
const userSqlImgDir = "/img/user_img/"

func saveUserImage(imageByte []byte) (bool, string) {
	return saveImage(imageByte, userImgDir, userSqlImgDir)
}

func PostDataDecrypt(postData string, dataChan chan string) {
	userPasswordByte, err := base64.StdEncoding.DecodeString(postData)
	if err != nil {
		log.Println("getUserPassword err")
		dataChan <- ""
	}
	userPasswordByte, err = PasswordHandler.RsaDecrypt(userPasswordByte)
	if err != nil {
		log.Println(err.Error())
		dataChan <- ""
	}
	dataChan <- string(userPasswordByte)
	close(dataChan)
}

func DoPasswordEncrypt(userName string, noEncryptPassword string) string {
	shaHashPostPassword := PasswordHandler.ShaHashData(userName, []byte(noEncryptPassword))
	aesHashPostPassword, err := PasswordHandler.AesEncrypt([]byte(shaHashPostPassword))
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	aesHashPostPasswordString := hex.EncodeToString(aesHashPostPassword)
	return aesHashPostPasswordString
}

func getUserIDFromUserName(userName string) int {
	sql := "select id from user_info where USER_NAME=?"
	var userID int
	err := DB.QueryRow(sql, userName).Scan(&userID)
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return userID
}

func getUserPasswordFromSql(userName string, sqlpasswordChannel chan string, userID chan int) {
	sql := "select id,USER_PASSWORD from user_info where USER_NAME=?"
	rows, err := DB.Query(sql, userName)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		sqlpasswordChannel <- ""
		userID <- -1
		return
	}
	var getPassword string
	var id int
	if rows.Next() {
		err := rows.Scan(&id, &getPassword)
		if err != nil {
			sqlpasswordChannel <- ""
			userID <- -1
			return
		}
	}
	userID <- id
	sqlpasswordChannel <- getPassword
}

func checkEncryptData(userID int, userName string, postData string, getSqlData string) (bool, error) {
	aesHashPostPasswordString := DoPasswordEncrypt(userName, postData)
	if len(aesHashPostPasswordString) == 0 || len(getSqlData) == 0 {
		return false, nil
	}
	return aesHashPostPasswordString == getSqlData, nil
}

func SelectUserName(id int64) (string, error) {
	sql := "select user_name from user_info where id=?"
	rows, err := DB.Query(sql, id)
	var name string = "ERROR"
	if err != nil {
		log.Println(err)
		return name, err
	}
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	defer rows.Close()
	return name, err
}

func Sign_In(userName string, password string) (int, error) {
	userIDChan := make(chan int)
	sqlPasswordChan := make(chan string)
	encryPostPassword := make(chan string)
	go PostDataDecrypt(password, encryPostPassword)
	go getUserPasswordFromSql(userName, sqlPasswordChan, userIDChan)
	userID := <-userIDChan
	result, err := checkEncryptData(userID, userName, <-encryPostPassword, <-sqlPasswordChan)
	if err == nil && result {
		log.Println("用户:" + userName + " 登陆成功")
		return userID, nil
	}
	log.Println(err.Error())
	log.Println("用户:" + userName + " 登陆失败")
	userID = -1
	return userID, nil
}

func CheckIsHadUser(username string) (bool, error) {
	sql := "SELECT COUNT(id) from user_info where USER_NAME=?"
	result := 0
	err := DB.QueryRow(sql, username).Scan(&result)
	if err != nil {
		log.Println(err.Error())
		return true, err
	}
	if result >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func saveRegistImage(isHadImage int, userImageBase64String string, imageUrlChan chan string) {
	var saveUrl string
	if isHadImage == 1 {
		imageByte, err := base64.StdEncoding.DecodeString(userImageBase64String)
		if err != nil {
			log.Println("图片base64转换失败")
		}
		isSaveFile, saveFileUrl := saveUserImage(imageByte)
		if !isSaveFile {
			log.Println("图片保存失败")
		}
		saveUrl = saveFileUrl
	} else {
		saveUrl = ""
	}
	imageUrlChan <- saveUrl
	close(imageUrlChan)
}

func RegistUser(username string, password string, findPasswordQuestion string, findPasswordAnswer string, userImage string, isHadImage int64) error {
	sql := "insert into user_info(USER_NAME,USER_PASSWORD,USER_PASSWORD_QUESTION,USER_PASSWORD_ANSWER,USER_IMAGE_URL,USER_IS_HAD_IMAGE) VALUES (?,?,?,?,?,?)"
	var saveUrlChan = make(chan string)
	var noEncrtyPassword = make(chan string)
	var noEncrtyFindPasswordAnswer = make(chan string)
	go saveRegistImage(int(isHadImage), userImage, saveUrlChan)
	go PostDataDecrypt(password, noEncrtyPassword)
	go PostDataDecrypt(password, noEncrtyFindPasswordAnswer)
	encryPassword := DoPasswordEncrypt(username, <-noEncrtyPassword)
	encryFindPasswordAnswer := DoPasswordEncrypt(username, <-noEncrtyFindPasswordAnswer)
	if len(encryPassword) == 0 || len(encryFindPasswordAnswer) == 0 {
		return errors.New("加密密码失败")
	}
	_, err := DB.Exec(sql, username, encryPassword, findPasswordQuestion, encryFindPasswordAnswer, <-saveUrlChan, isHadImage)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func FindPasswordQuestion(username string) (string, error) {
	sql := "SELECT USER_PASSWORD_QUESTION from user_info where USER_NAME=?"
	var question string
	err := DB.QueryRow(sql, username).Scan(&question)
	if err != nil {
		return "", err
	} else {
		return question, nil
	}
}

func CheckQuestionAnswer(username string, answer string) (string, error) {
	sql := "select USER_PASSWORD_ANSWER from user_info where USER_NAME=?"
	var sqlPasswordAnswer string
	decryAnswer := make(chan string)
	go PostDataDecrypt(answer, decryAnswer)
	err := DB.QueryRow(sql, username).Scan(&sqlPasswordAnswer)
	if err != nil {
		return ERROR, err
	}
	userID := getUserIDFromUserName(username)
	if userID == -1 || userID == 0 {
		return ERROR, nil
	}
	result, _ := checkEncryptData(userID, username, <-decryAnswer, sqlPasswordAnswer)
	if result {
		return SUCCESS, nil
	} else {
		return ERROR, nil
	}
}
func UpdateUserPassword(username string, password string) (string, error) {
	sql := "update user_info set USER_PASSWORD=? where USER_NAME=?"
	var noDecryptPassword = make(chan string)
	go PostDataDecrypt(password, noDecryptPassword)
	decryPassword := DoPasswordEncrypt(username, <-noDecryptPassword)
	_, err := DB.Exec(sql, decryPassword, username)
	if err != nil {
		log.Println(err.Error())
		return ERROR, err
	} else {
		return SUCCESS, nil
	}
}
