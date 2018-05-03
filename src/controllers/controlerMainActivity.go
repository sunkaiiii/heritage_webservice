package controllers

import (
	"log"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func GetMainDivideActivityImageUrl(c *gin.Context) {
	c.String(200, models.GetMainDivideActivityImageUrl())
}

func GetChannelInformation(c *gin.Context) {
	divide := c.Query("divide")
	c.String(200, models.GetChannelInformation(divide))
}

func GetChannelFolkInformation(c *gin.Context) {
	c.String(200, models.GetChannelFolkInformation())
}

func GetChannelFolkSingleInformation(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(200, ERROR)
		return
	}
	log.Println("获取" + strconv.Itoa(id) + "的channelFolk数据")
	c.String(200, models.GetChannelFolkSingleInformation(id))
}

func SearchChannelForkInfo(c *gin.Context) {
	searchInfo := c.Query("searchInfo")
	if searchInfo == "" {
		c.String(200, ERROR)
		return
	}
	log.Println("搜索:" + searchInfo)
	c.String(200, models.SearchChannelForkInfo(searchInfo))
}

func SearchBottomNewsInformation(c *gin.Context) {
	searchInfo := c.Query("searchInfo")
	if searchInfo == "" {
		c.String(200, ERROR)
		return
	}
	log.Println("搜索聚焦非遗：" + searchInfo)
	c.String(200, models.SearchBottomNewsInformation(searchInfo))
}

func SearchFolkNewsInformaiton(c *gin.Context) {
	searchInfo := c.Query("searchInfo")
	if searchInfo == "" {
		c.String(200, ERROR)
		return
	}
	log.Println("搜索全部新闻：" + searchInfo)
	c.String(200, models.SearchFolkNewsInformaiton(searchInfo))
}
