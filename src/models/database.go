package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var RedisDB *redis.Client

const ERROR = "ERROR"
const SUCCESS = "SUCCESS"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/heritage")
	if err != nil {
		log.Fatalf("Open database error:%s\n", err)
	}
	err2 := db.Ping()
	if err2 != nil {
		log.Fatal(err2)
	}
	DB = db
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	pong, err := client.Ping().Result()
	log.Println(pong)
	if err != nil {
		log.Fatal(err)
	}
	RedisDB = client
	return db, err
}

func UpdatePassword() {
	sql := "select ID,USER_PASSWORD from user_info"
	rows, _ := DB.Query(sql)
	defer rows.Close()
	var id int
	var password string
	for rows.Next() {
		rows.Scan(&id, &password)
		passwordByte, _ := base64.StdEncoding.DecodeString(password)
		encryptPassword, _ := RsaDecrypt(passwordByte)
		shaPassword := shaHashData(encryptPassword)
		sql2 := "update user_info set USER_PASSWORD=? where ID=? "
		DB.Exec(sql2, shaPassword, id)
	}
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	file, _ := os.Open("./models/public.pem")
	buf := make([]byte, 1024)
	file.Read(buf)
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	file, err := os.Open("./models/private.pem")
	if err != nil {
		log.Println("cannot open PrivateKey.pem")
		return nil, err
	}
	buf := make([]byte, 1024)
	file.Read(buf)
	block, _ := pem.Decode(buf)
	if block == nil {
		println("private key error")
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		println("error")
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func saveImage(imageByte []byte, imgDir string, sqlImgDir string) (bool, string) {
	imgtime := strconv.FormatInt(time.Now().UnixNano(), 10)
	imgName := imgDir + imgtime + ".jpg"
	saveImg, err := os.Create(imgName)
	if err != nil {
		log.Println(err.Error())
		return false, ERROR
	}
	commentImageURL := sqlImgDir + imgtime + ".jpg"
	_, err = saveImg.Write(imageByte)
	return true, commentImageURL
}
