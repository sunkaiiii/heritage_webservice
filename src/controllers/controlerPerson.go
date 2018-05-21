package controllers

import (
	"log"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func GetFollowNumber(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, strconv.Itoa(models.GetFollowNumber(userID)))
}

func GetFansNumber(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, strconv.Itoa(models.GetFansNumber(userID)))
}

func GetUserPermission(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, strconv.Itoa(models.GetUserPermission(userID)))
}

func GetUserFocusAndFansViewPermission(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, strconv.Itoa(models.GetUserFocusAndFansViewPermission(userID)))
}

func SetUserPermission(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	permission, err := strconv.Atoi(c.Query("permission"))
	if err != nil {
		log.Println("permission有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.SetUserPermission(userID, permission))
}

func SetUserFocusAndFansViewPermission(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	permission, err := strconv.Atoi(c.Query("permission"))
	if err != nil {
		log.Println("permission有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.SetUserFocusAndFansViewPermission(userID, permission))
}

func GetUserAllInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserAllInfo(userID))
}

func GetFollowInformation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetFollowInformation(userID))
}

func GetFansInformation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetFansInformation(userID))
}

func AddFocus(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	focusID, err := strconv.Atoi(c.Query("focusID"))
	if err != nil {
		log.Println("focusID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.AddFocus(userID, focusID))
}

func CancelFocus(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	focusID, err := strconv.Atoi(c.Query("focusID"))
	if err != nil {
		log.Println("focusID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.CancelFocus(userID, focusID))
}

func CheckFollowEachother(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	focusID, err := strconv.Atoi(c.Query("focusID"))
	if err != nil {
		log.Println("focusID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.CheckFollowEachother(userID, focusID))
}

func GetSearchUserInfo(c *gin.Context) {
	searchName := c.Query("searchName")
	c.String(200, models.GetSearchUserInfo(searchName))
}

func GetUserImage(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetUserImage(userID))
}

func UpdateUserImage(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	imageString := c.PostForm("image")
	c.String(200, models.UpdateUserImage(userID, imageString))
}

func AddUserCollection(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	typeID, err := strconv.Atoi(c.Query("typeID"))
	if err != nil {
		log.Println("typeID")
		c.String(200, ERROR)
		return
	}
	collectType := c.Query("type")
	c.String(200, models.AddUserCollection(userID, collectType, typeID))
}
func CancelUserCollect(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	typeID, err := strconv.Atoi(c.Query("typeID"))
	if err != nil {
		log.Println("typeID")
		c.String(200, ERROR)
		return
	}
	collectType := c.Query("type")
	c.String(200, models.CancelUserCollect(userID, collectType, typeID))
}
func GetUserCollection(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	collectType := c.Query("type")
	log.Println("用户:" + strconv.Itoa(userID) + "查询collect" + collectType)
	c.String(200, models.GetUserCollection(userID, collectType))
}

func CheckIsCollection(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		log.Println("userID有误")
		c.String(200, ERROR)
		return
	}
	typeID, err := strconv.Atoi(c.Query("typeID"))
	if err != nil {
		log.Println("typeID")
		c.String(200, ERROR)
		return
	}
	collectType := c.Query("type")
	c.String(200, models.CheckIsCollection(userID, collectType, typeID))
}
