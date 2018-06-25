package models

import (
	"log"
	"os"
	"strconv"
	"time"
)

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
