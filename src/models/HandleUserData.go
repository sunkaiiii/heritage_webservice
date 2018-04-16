package models

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"
)

const userImgDir = "./img/user_img/"
const userSqlImgDir = "/img/user_img/"

func saveUserImage(imageByte []byte) (bool, string) {
	return saveImage(imageByte, userImgDir, userSqlImgDir)
}

func PostDataEncrypt(postData string, dataChan chan string) {
	userPasswordByte, err := base64.StdEncoding.DecodeString(postData)
	if err != nil {
		log.Println("userPassWordByte err")
		dataChan <- ""
	}
	userPasswordByte, err = RsaDecrypt(userPasswordByte)
	if err != nil {
		log.Println("userPasswordByte2 err")
		dataChan <- ""
	}
	userPassword := base64.StdEncoding.EncodeToString(userPasswordByte)
	dataChan <- userPassword
}

func sqlDataEncrypt(getSqlData string, dataChan chan string) {
	sqlByte, err := base64.StdEncoding.DecodeString(getSqlData)
	if err != nil {
		log.Println("sqlByte err")
		dataChan <- ""
	}
	sqlByte, err = RsaDecrypt(sqlByte)
	if err != nil {
		log.Println("sqlByte2 err")
		dataChan <- ""
	}
	sqlPassword := base64.StdEncoding.EncodeToString(sqlByte)
	dataChan <- sqlPassword
}

func checkEncryptData(postData string, getSqlData string) (bool, error) {
	getSqlData = strings.Replace(getSqlData, "\n", "", -1)
	sqlChan := make(chan string)
	userChan := make(chan string)
	go sqlDataEncrypt(getSqlData, sqlChan)
	go PostDataEncrypt(postData, userChan)
	sqlPassword := <-sqlChan
	userPassword := <-userChan
	log.Println(sqlPassword)
	log.Println(userPassword)
	return sqlPassword == userPassword, nil
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
	sql := "select id,USER_PASSWORD from user_info where USER_NAME=?"
	rows, err := DB.Query(sql, userName)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	var getPassword string
	var id int
	if rows.Next() {
		err := rows.Scan(&id, &getPassword)
		if err != nil {
			return -1, err
		}
	}
	defer rows.Close()
	result, err := checkEncryptData(password, getPassword)
	if err == nil && result {
		return id, nil
	}
	return -1, nil
}

func CheckIsHadUser(username string) (bool, error) {
	sql := "SELECT USER_NAME from USER_INFO where USER_NAME=?"
	rows, err := DB.Query(sql, username)
	defer rows.Close()
	if err != nil {
		return true, err
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func RegistUser(username string, password string, findPasswordQuestion string, findPasswordAnswer string, userImage string, isHadImage int64) error {
	sql := "insert into USER_INFO(USER_NAME,USER_PASSWORD,USER_PASSWORD_QUESTION,USER_PASSWORD_ANSWER,USER_IMAGE_URL,USER_IS_HAD_IMAGE) VALUES (?,?,?,?,?,?)"
	var saveUrl string
	if isHadImage == 1 {
		imageByte, err := base64.StdEncoding.DecodeString(userImage)
		if err != nil {
			return err
		}
		isSaveFile, saveFileUrl := saveUserImage(imageByte)
		if !isSaveFile {
			log.Println("图片保存失败")
			return errors.New("图片保存失败")
		}
		saveUrl = saveFileUrl
	} else {
		saveUrl = ""
	}
	_, err := DB.Exec(sql, username, password, findPasswordQuestion, findPasswordAnswer, saveUrl, isHadImage)
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
	err := DB.QueryRow(sql, username).Scan(&sqlPasswordAnswer)
	if err != nil {
		return ERROR, err
	}
	result, _ := checkEncryptData(answer, sqlPasswordAnswer)
	if result {
		return SUCCESS, nil
	} else {
		return ERROR, nil
	}
}
func UpdateUserPassword(username string, password string) (string, error) {
	sql := "update user_info set USER_PASSWORD=? where USER_NAME=?"
	_, err := DB.Exec(sql, password, username)
	if err != nil {
		log.Println(err.Error())
		return ERROR, err
	} else {
		return SUCCESS, nil
	}
}
