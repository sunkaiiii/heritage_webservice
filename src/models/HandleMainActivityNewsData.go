package models

import (
	"encoding/json"
	"log"
	"strconv"
)

func GetFolkNewsList(category string, start int, end int) string {
	key := "MainActivityNews" + category + strconv.Itoa(start) + "_" + strconv.Itoa(end)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,title,time,category,content,img from folk_news where category=? LIMIT ?,?"
	rows, err := DB.Query(sql, category, start, (end - start))
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultArray := make([]FolkNewsLite, end-start)
	count := 0
	for rows.Next() {
		var data FolkNewsLite
		err = rows.Scan(&data.ID, &data.Title, &data.Time, &data.Category, &data.Content, &data.Img)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		resultArray[count] = data
		count++
	}
	jsonResult, err := json.Marshal(resultArray[0:count])
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	RedisDB.Set(key, jsonString, 0)
	return jsonString
}

func getFolkNewsDataClass(id int) (FolkNewsLite, error) {
	sql := "select id,title,time,category,content,img from folk_news where id=?"
	var data FolkNewsLite
	err := DB.QueryRow(sql, id).Scan(&data.ID, &data.Title, &data.Time, &data.Category, &data.Content, &data.Img)
	return data, err
}

func GetFolkNewsInformation(id int) string {
	key := "folknewsInformation" + "_" + strconv.Itoa(id)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select details from folk_news where id=?"
	var data string
	err = DB.QueryRow(sql, id).Scan(&data)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	RedisDB.Set(key, data, 0)
	return data
}
func getBottomNewsInformationClassByID(id int) (BottomNewsLite, error) {
	sql := "select id,title,time,news_briefly,img from bottom_folk_news where id=?"
	var data BottomNewsLite
	err := DB.QueryRow(sql, id).Scan(&data.ID, &data.Title, &data.Time, &data.Briefly, &data.Img)
	return data, err
}
func GetBottomNewsLiteInformation(start int, end int) string {
	key := "bottomNewsLiteInfor_" + strconv.Itoa(start) + "_" + strconv.Itoa(end)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,title,time,news_briefly,img from bottom_folk_news LIMIT ?,?"
	count := end - start
	rows, err := DB.Query(sql, start, count)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]BottomNewsLite, count)
	for i := 0; rows.Next(); i++ {
		var data BottomNewsLite
		err = rows.Scan(&data.ID, &data.Title, &data.Time, &data.Briefly, &data.Img)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		resultList[i] = data
	}
	jsonResult, err := json.Marshal(resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	RedisDB.Set(key, jsonString, 0)
	return string(jsonString)
}

func GetBottomNewsInformationByID(id int) string {
	key := "getBottomNewsInformationById_" + strconv.Itoa(id)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,title,time,news_briefly,img,content from bottom_folk_news where id=?"
	var data BottomNews
	err = DB.QueryRow(sql, id).Scan(&data.ID, &data.Title, &data.Time, &data.Briefly, &data.Img, &data.Content)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonResult, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	RedisDB.Set(key, jsonString, 0)
	return jsonString
}

func GetMainPageSlideNewsInformation() string {
	key := "MainPageSlideNews"
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,content,img,detail from main_page_slide LIMIT 0,5"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]MainPageSlideNews, 5)
	for i := 0; rows.Next(); i++ {
		var data MainPageSlideNews
		err = rows.Scan(&data.ID, &data.Content, &data.Img, &data.Detail)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		resultList[i] = data
	}
	jsonResult, err := json.Marshal(resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	RedisDB.Set(key, jsonString, 0)
	return jsonString
}
