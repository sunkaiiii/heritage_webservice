package models

import (
	"log"
	"strconv"
)

func GetMainDivideActivityImageUrl() string {
	key := "MainDivideActivityImageUrl"
	result, err := GetRedisKey(key)
	if err == nil {
		return result
	}
	sql := "select id,category,url from main_activity_divide_img"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		return ERROR
	}
	resultArray := make([](*ActivityDivideImage), 5, 20)
	count := 0
	for rows.Next() {
		var data ActivityDivideImage
		err = rows.Scan(&data.Id, &data.Category, &data.Url)
		if err != nil {
			return ERROR
		}
		resultArray[count] = &data
		count++
	}
	jsonString := masharlData(resultArray[0:count])
	SetRedisKey(key, jsonString)
	return jsonString
}

func getChannelInformationSingleClass(id int) (ChannelInformaiton, error) {
	sql := "SELECT id,time,category,location,apply_location,content,number,img,title from classify_activity_new where id=?"
	var data ChannelInformaiton
	err := DB.QueryRow(sql, id).Scan(&data.Id, &data.Time, &data.Category, &data.Location, &data.Apply_location, &data.Content, &data.Number, &data.Img, &data.Title)
	return data, err
}

func GetChannelInformation(divide string) string {
	key := "Channel" + divide
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "SELECT id,time,category,location,apply_location,content,number,img,title from classify_activity_new where divide=? and LENGTH (img)>0"
	rows, err := DB.Query(sql, divide)
	defer rows.Close()
	if err != nil {
		return ERROR
	}
	var resultArray [](*ChannelInformaiton)
	for rows.Next() {
		var data ChannelInformaiton
		err = rows.Scan(&data.Id, &data.Time, &data.Category, &data.Location, &data.Apply_location, &data.Content, &data.Number, &data.Img, &data.Title)
		if err != nil {
			return ERROR
		}
		resultArray = append(resultArray, &data)
	}
	jsonString := masharlData(resultArray)
	SetRedisKey(key, jsonString)
	return jsonString
}

func GetChannelFolkInformation() string {
	key := "ChannelFolkInformation"
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "SELECT id,divide,title,apply_location,img,category from folk_activity_information where LENGTH(img)>0"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	var resultArray [](*ChannelForkInformationLite)
	for rows.Next() {
		var data ChannelForkInformationLite
		err = rows.Scan(&data.Id, &data.Divide, &data.Title, &data.Apply_location, &data.Img, &data.Category)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		resultArray = append(resultArray, &data)
	}
	jsonString := masharlData(resultArray)
	SetRedisKey(key, jsonString)
	return jsonString

}

func GetChannelFolkSingleInformation(id int) string {
	key := "ChannelFolkSingleInformation_" + strconv.Itoa(id)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "SELECT id,time,divide,category,location,title,apply_location,content,number,img from folk_activity_information where id=?"
	var data FolkData
	err = DB.QueryRow(sql, id).Scan(&data.Id, &data.Time, &data.Divide, &data.Category, &data.Location, &data.Title, &data.Apply_location, &data.Content, &data.Number, &data.Img)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := masharlData(data)
	SetRedisKey(key, jsonString)
	return jsonString
}

func SearchChannelForkInfo(searchInfo string) string {
	key := "ChannelFolkInfo" + searchInfo
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,divide,title,apply_location,img,category from folk_activity_information where (title like ? or divide like ? or location like ?  or category like ?) and LENGTH(img)>0"
	searchInfo = "%" + searchInfo + "%"
	rows, err := DB.Query(sql, searchInfo, searchInfo, searchInfo, searchInfo)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	var resultArray [](*ChannelForkInformationLite)
	for rows.Next() {
		var data ChannelForkInformationLite
		err = rows.Scan(&data.Id, &data.Divide, &data.Title, &data.Apply_location, &data.Img, &data.Category)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		resultArray = append(resultArray, &data)
	}
	jsonString := masharlData(resultArray)
	SetRedisKey(key, jsonString)
	return jsonString
}
