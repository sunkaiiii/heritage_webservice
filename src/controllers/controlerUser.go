package controllers

import (
	"log"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

const ERROR = "ERROR"
const SUCCESS = "SUCCESS"

func GetUserName(c *gin.Context) {
	id := c.Query("id")
	nid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	name, err := models.SelectUserName(nid)
	if err != nil {
		c.String(200, "ERROR")
	}
	c.JSON(200, name)
}
func Sign_In(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	id, _ := models.Sign_In(username, password)
	if id > 0 {
		log.Println("用户:" + username + " 登陆")
	}
	c.String(200, strconv.Itoa(id))
}

func UserRegist(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	findPasswordQuestion := c.PostForm("findPasswordQuestion")
	findPasswordAnswer := c.PostForm("findPasswordAnswer")
	userImage := c.PostForm("userImage")
	if checkIsHadUser(username) {
		log.Println("已有该用户:" + username)
		c.String(200, "0")
		return
	}
	var isHadImage int64
	if userImage != "" {
		log.Println(username + " 有图")
		isHadImage = 1
	} else {
		isHadImage = 0
	}
	err := models.RegistUser(username, password, findPasswordQuestion, findPasswordAnswer, userImage, isHadImage)
	if err != nil {
		c.String(200, ERROR)
		return
	}
	log.Println("用户：" + username + " 注册成功")
	c.String(200, "1")
	return
}

func checkIsHadUser(username string) bool {
	result, err := models.CheckIsHadUser(username)
	if err != nil {
		return true
	}
	return result
}

func FindPassWordQuestion(c *gin.Context) {
	username := c.Query("username")
	result, err := models.FindPasswordQuestion(username)
	if err == nil {
		log.Println("用户:" + username + " 查询问题:" + result)
		c.String(200, result)
	} else {
		c.String(200, ERROR)
	}
	return
}

func CheckQuestionAnswer(c *gin.Context) {
	username := c.PostForm("username")
	answer := c.PostForm("answer")
	result, err := models.CheckQuestionAnswer(username, answer)
	if err != nil || result == ERROR {
		c.String(200, ERROR)
	} else {
		c.String(200, result)
	}
	return
}

func ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	result, err := models.UpdateUserPassword(username, password)
	if err != nil || result == ERROR {
		c.String(200, ERROR)
	} else {
		c.String(200, SUCCESS)
	}
}
