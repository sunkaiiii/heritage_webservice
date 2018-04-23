package controllers

import (
	"log"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func GetFindActivityID(c *gin.Context) {
	c.String(200, models.GetFindActivityID())
	return
}

func AddUserCommentInformation(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("userID"))
	if err != nil {
		c.String(200, ERROR)
		return
	}
	commentTitle := c.PostForm("commentTitle")
	commentContent := c.PostForm("commentContent")
	commentImage := c.PostForm("commentImage")
	location := c.PostForm("location")
	result := models.AddUserCommentInformation(userID, commentTitle, commentContent, commentImage, location)
	if SUCCESS == result {
		log.Println("用户：" + strconv.Itoa(userID) + "添加了comment " + commentTitle)
	} else {
		log.Println("用户：" + strconv.Itoa(userID) + "添加了comment 失败")
	}
	c.String(200, result)
}

func IsUserFollow(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.String(200, ERROR)
		return
	}
	fansID, err := strconv.Atoi(c.Query("fansID"))
	if err != nil {
		c.String(200, ERROR)
		return
	}
	c.String(200, models.IsUserFollow(userID, fansID))
}

func GetUserCommentInformation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.String(200, ERROR)
		return
	}
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		c.String(200, ERROR)
		return
	}
	log.Println("用户:" + strconv.Itoa(userID) + "查看发现页内容")
	c.String(200, models.GetUserCommentInformation(userID, start))
}

func DeleteUserCommentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.DeleteUserCommentByID(id))
}

func GetUserCommentImageUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserCommentImageUrl(id))
}

func UpdateUserCommentImage(c *gin.Context) {
	idString := c.DefaultPostForm("id", strconv.Itoa(-1))
	imgString := c.PostForm("image")
	id, err := strconv.Atoi(idString)
	if err != nil || id == -1 {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.UpdateUserCommentImage(id, imgString))
}

func UpdateUserCommentInformaiton(c *gin.Context) {
	commentIDString := c.PostForm("id")
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	commentTitle := c.PostForm("title")
	commentContent := c.PostForm("content")
	commentImg := c.PostForm("image")
	log.Println("更新" + strconv.Itoa(commentID) + "内容")
	c.String(200, models.UpdateUserCommentInformaiton(commentID, commentTitle, commentContent, commentImg))
}

func GetUserCommentIdByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserCommentIdByUser(userID))
}

func GetUserCommentInformationByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		log.Println("start有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserCommentInformationByUser(userID, start))
}

func GetUserCommentInformaitonByOwn(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		log.Println("start有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserCommentInformaitonByOwn(userID, start))
}

func GetUserCommentInformationBySameLocation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("id有误")
		c.String(200, ERROR)
		return
	}
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		log.Println("start有误")
		c.String(200, ERROR)
		return
	}
	location := c.Query("location")
	c.String(200, models.GetUserCommentInformationBySameLocation(userID, start, location))
}

func GetAllUserCommentInfoByID(c *gin.Context) {
	user, err := strconv.Atoi(c.Query("user"))
	if err != nil {
		log.Println("user有误")
		c.String(200, ERROR)
		return
	}
	commentID, err := strconv.Atoi(c.Query("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetAllUserCommentInfoByID(user, commentID))
}

func GetCommentLikeNumber(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Query("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, strconv.Itoa(models.GetCommentLikeNumber(commentID)))
}

func GetUserCommentCount(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Query("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	minireply, err := strconv.Atoi(c.Query("miniReply"))
	if err != nil {
		log.Println("minireply有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, strconv.Itoa(models.GetUserCommentCount(commentID, minireply)))
}

func GetUserCommentReply(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Query("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	minireply, err := strconv.Atoi(c.Query("miniReply"))
	if err != nil {
		log.Println("miniReply有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserCommentReply(commentID, minireply))
}

func SetUserLike(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	commentID, err := strconv.Atoi(c.Query("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.SetUserLike(userID, commentID))
}

func CancelUserLike(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	commentID, err := strconv.Atoi(c.Query("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.CancelUserLike(userID, commentID))
}

func AddUserCommentReply(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	commentID, err := strconv.Atoi(c.PostForm("commentID"))
	if err != nil {
		log.Println("commentID有误")
		c.String(200, ERROR)
		return
	}
	reply := c.PostForm("reply")
	c.String(200, models.AddUserCommentReply(userID, commentID, reply))
}

func UpdateUserCommentReply(c *gin.Context) {
	replyID, err := strconv.Atoi(c.PostForm("replyID"))
	if err != nil {
		log.Println("replyID有误")
		c.String(200, ERROR)
		return
	}
	reply := c.PostForm("reply")
	c.String(200, models.UpdateUserCommentReply(replyID, reply))
}

func DeleteUserCommentReply(c *gin.Context) {
	replyID, err := strconv.Atoi(c.Query("replyID"))
	if err != nil {
		log.Println("replyID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.DeleteUserCommentReply(replyID))
}

func GetUserLikeComment(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserLikeComment(userID))
}
