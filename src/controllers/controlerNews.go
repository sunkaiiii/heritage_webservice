package controllers

import (
	"log"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func GetFolkNewsList(c *gin.Context) {
	divide := c.Query("divide")
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		log.Println(err.Error())
		c.String(200, ERROR)
		return
	}
	end, err := strconv.Atoi(c.Query("end"))
	if err != nil {
		log.Println(err.Error())
		c.String(200, ERROR)
	}

	c.String(200, models.GetFolkNewsList(divide, start, end))
}

func GetFolkNewsInformation(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err.Error())
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetFolkNewsInformation(id))
}

func GetBottomNewsLiteInformation(c *gin.Context) {
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		log.Println(err.Error())
		c.String(200, ERROR)
		return
	}
	end, err := strconv.Atoi(c.Query("end"))
	if err != nil {
		log.Println(err.Error())
		c.String(200, ERROR)
	}
	c.String(200, models.GetBottomNewsLiteInformation(start, end))
}

func GetBottomNewsInformationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err.Error())
		c.String(200, ERROR)
		return
	}
	c.String(200, models.GetBottomNewsInformationByID(id))
}

func GetMainPageSlideNewsInformation(c *gin.Context){
	c.String(200,models.GetMainPageSlideNewsInformation())
}
