package models

import (
	"database/sql"
	"encoding/base64"
	"log"

	PasswordHandler "../password"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var RedisDB *redis.Client

const ERROR = "ERROR"
const SUCCESS = "SUCCESS"

func init() {
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
}
func InitDB() *sql.DB {
	return DB
}

func UpdatePassword() {
	sql := "select ID,USER_NAME,USER_PASSWORD from user_info"
	rows, _ := DB.Query(sql)
	defer rows.Close()
	var id int
	var username
	var password string
	for rows.Next() {
		rows.Scan(&id,&username, &password)
		passwordByte, _ := base64.StdEncoding.DecodeString(password)
		encryptPassword, _ := PasswordHandler.RsaDecrypt(passwordByte)
		shaPassword := PasswordHandler.ShaHashData(id,username,ncryptPassword)
		sql2 := "update user_info set USER_PASSWORD=? where ID=? "
		DB.Exec(sql2, shaPassword, id)
	}
}
