package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
)

func getFolkNewsData(rows *sql.Rows, resultList *[]FolkNewsLite) (int, error) {
	count := 0
	for rows.Next() {
		var data FolkNewsLite
		err := rows.Scan(&data.ID, &data.Title, &data.Time, &data.Category, &data.Content, &data.Img)
		if err != nil {
			return -1, err
		}
		(*resultList)[count] = data
		count++
	}
	return count, nil
}

func GetFolkNewsList(category string, start int, count int) string {
	key := "MainActivityNews" + category + strconv.Itoa(start) + "_" + strconv.Itoa(count)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,title,time,category,content,img from folk_news where category=? LIMIT ?,?"
	rows, err := DB.Query(sql, category, start, count)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]FolkNewsLite, count)
	count, err = getFolkNewsData(rows, &resultList)
	jsonResult := masharlData(resultList[0:count])
	if jsonResult == ERROR {
		return ERROR
	}
	SetRedisKey(key, jsonResult)
	return jsonResult
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

func getBottomNewsData(rows *sql.Rows, resultList *[]BottomNewsLite) (int, error) {
	count := 0
	for i := 0; rows.Next(); i++ {
		var data BottomNewsLite
		err := rows.Scan(&data.ID, &data.Title, &data.Time, &data.Briefly, &data.Img)
		if err != nil {
			log.Println(err.Error())
			return count, err
		}
		(*resultList)[count] = data
		count++
	}
	return count, nil
}

func GetBottomNewsLiteInformation(start int, count int) string {
	key := "bottomNewsLiteInfor_" + strconv.Itoa(start) + "_" + strconv.Itoa(count)
	result, err := RedisDB.Get(key).Result()
	if err == nil {
		return result
	}
	sql := "select id,title,time,news_briefly,img from bottom_folk_news LIMIT ?,?"
	rows, err := DB.Query(sql, start, count)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]BottomNewsLite, count)
	count, err = getBottomNewsData(rows, &resultList)
	jsonString := masharlData(resultList[0:count])
	RedisDB.Set(key, jsonString, 0)
	return jsonString
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

func SearchBottomNewsInformation(searchString string) string {
	key := "SearchBottomNewsInformation_" + searchString
	result, err := GetRedisKey(key)
	if err == nil {
		return result
	}
	sql := "select id,title,time,news_briefly,img from bottom_folk_news where title like ? or news_briefly like ? LIMIT 0,50"
	searchString = "%" + searchString + "%"
	rows, err := DB.Query(sql, searchString, searchString)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]BottomNewsLite, 50)
	count, err := getBottomNewsData(rows, &resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := masharlData(resultList[0:count])
	if ERROR != jsonString {
		SetRedisKey(key, jsonString)
	}
	return jsonString
}

func SearchFolkNewsInformaiton(searhString string) string {
	key := "SearhchFolkNewsInformation_" + searhString
	result, err := GetRedisKey(key)
	if err == nil {
		return result
	}
	sql := "select id,title,time,category,content,img from folk_news where title like ? LIMIT 0,50"
	searhString = "%" + searhString + "%"
	rows, err := DB.Query(sql, searhString)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	defer rows.Close()
	resultList := make([]FolkNewsLite, 50)
	count, err := getFolkNewsData(rows, &resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonResult := masharlData(resultList[0:count])
	if jsonResult != ERROR {
		SetRedisKey(key, jsonResult)
	}
	return jsonResult
}
