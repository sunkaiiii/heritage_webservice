package controllers

import (
	"log"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func AddPushMessage(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("userID"))
	if err != nil {
		log.Println("userID有误 " + err.Error())
		c.String(200, ERROR)
		return
	}
	replyCommentID, err := strconv.Atoi(c.PostForm("replyCommentID"))
	if err != nil {
		log.Println("replyCommentID " + err.Error())
		c.String(200, ERROR)
		return
	}
	replyToUserID, err := strconv.Atoi(c.PostForm("replyToUserID"))
	if err != nil {
		log.Println("replyToUserID " + err.Error())
		c.String(200, ERROR)
		return
	}
	userName := c.PostForm("userName")
	replyContent := c.PostForm("replyContent")
	replyToUserName := c.PostForm("replyToUserName")
	replyTime := c.PostForm("replyTime")
	originalReplyContent := c.PostForm("originalReplyContent")
	c.String(200, models.AddPushMessage(userID, userName, replyCommentID, replyContent, replyToUserID, replyToUserName, replyTime, originalReplyContent))
}

func GetPushMessage(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误 " + err.Error())
		c.String(200, ERROR)
		return
	}
	c.String(200, models.SendSinglePushMessageInfo(userID))
}

func GetAllPushReplyInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误 " + err.Error())
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetPushMessageInfo(userID))
}
