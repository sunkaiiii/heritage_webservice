package models

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"time"
)

const findImgDir = "./img/find_img/"
const findSqlImgDir = "/img/find_img/"

const miniReply = -1
const normalReply = 0

func SaveFindImage(imageByte []byte) (bool, string) {
	return saveImage(imageByte, findImgDir, findSqlImgDir)
}

func GetFindActivityID() string {
	key := "GetFindActivityID"
	result, err := GetRedisKey(key)
	if err == nil {
		return result
	}
	sql := "select id from find_activity order by id desc LIMIT 0,4"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	dataArray := make([]int, 4)
	count := 0
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		dataArray[count] = id
		count++
	}
	jsonResult, err := json.Marshal(dataArray)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	SetRedisKey(key, jsonString)
	return jsonString
}

func GetFindActivityInformation(id int) string {
	key := "GetFindActivityInformation" + strconv.Itoa(id)
	result, err := GetRedisKey(key)
	if err == nil {
		return result
	}
	sql := "select title,content,image from find_activity where id=?"
	var data FindActivityData
	var imageByte []byte
	err = DB.QueryRow(sql, id).Scan(&data.Title, &data.Content, &imageByte)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	imageBase64 := base64.StdEncoding.EncodeToString(imageByte)
	data.Image = imageBase64
	jsonResult, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	SetRedisKey(key, result)
	return string(jsonString)
}

func AddUserCommentInformation(userID int, commentTitle string, commentContent string, commentImage string) string {
	sql := "insert into user_comment(user_id,comment_time,comment_title,comment_content,comment_image_url) VALUES (?,?,?,?,?)"
	imageByte, err := base64.StdEncoding.DecodeString(commentImage)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	isSave, commentImageUrl := SaveFindImage(imageByte)
	if !isSave {
		log.Println("图片保存错误")
		return ERROR
	}
	commentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = DB.Exec(sql, userID, commentTime, commentTitle, commentContent, commentImageUrl)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	return SUCCESS
}

func IsUserFollow(userID int, fansID int) string {
	sql := "select * from my_focus where focus_userID=? and focus_fansID=?"
	rows, err := DB.Query(sql, userID, fansID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	if rows.Next() {
		return SUCCESS
	}
	return ERROR
}

func GetUserIsLike(userID int, commentID int) string {
	sql := "SELECT * from user_comment_like where userID=? and commentID=?"
	rows, err := DB.Query(sql, userID, commentID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	if rows.Next() {
		return SUCCESS
	}
	return ERROR
}

func GetUserCommentInformation(userID int, start int) string {
	sql := "SELECT  user_comment.id,user_id,user_name,comment_time,comment_title,comment_content,comment_image_url " +
		"from user_comment,user_info where user_id=user_info.ID order by id DESC LIMIT ?,20"
	rows, err := DB.Query(sql, start)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	count := 0
	resultArray := make([]UserCommentData, 20)
	for rows.Next() {
		var data UserCommentData
		err = rows.Scan(&data.ID, &data.UserID, &data.UserName, &data.CommentTime, &data.CommentTitle, &data.ComemntContent, &data.ImageURL)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		data.IsFollow = IsUserFollow(userID, data.UserID)
		data.IsLike = GetUserIsLike(userID, data.ID)
		sql = "SELECT COUNT(*) from user_comment_like where commentID=?"
		err = DB.QueryRow(sql, data.ID).Scan(&data.LikeNum)
		if err != nil {
			log.Println(err.Error())
			data.LikeNum = 0
		}
		sql = "SELECT COUNT(*) from user_comment_reply where comment_ID=?"
		err = DB.QueryRow(sql, data.ID).Scan(&data.ReplyNum)
		if err != nil {
			log.Println(err.Error())
			data.ReplyNum = 0
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
	return jsonString
}

func deleteUserCommentReplyByCommentIDChan(id int, deleteCommentInfo chan string) {
	sql := "delete from user_comment_reply where comment_ID=?"
	_, err := DB.Exec(sql, id)
	if err != nil {
		deleteCommentInfo <- ERROR
	} else {
		deleteCommentInfo <- SUCCESS
	}
}

func deleteFromUserCommentLikeByCommentIDChan(id int, deleteCommentLikeInfo chan string) {
	sql := "delete from user_comment_like where commentID=?"
	_, err := DB.Exec(sql, id)
	if err != nil {
		deleteCommentLikeInfo <- ERROR
	} else {
		deleteCommentLikeInfo <- SUCCESS
	}
}

func DeleteUserCommentByID(id int) string {
	sql := "delete from user_comment where id=?"
	deleteCommentInfo := make(chan string)
	deleteCommentLikeInfo := make(chan string)
	go deleteUserCommentReplyByCommentIDChan(id, deleteCommentInfo)
	go deleteFromUserCommentLikeByCommentIDChan(id, deleteCommentLikeInfo)
	_, err := DB.Exec(sql, id)
	commentInfo := <-deleteCommentInfo
	likeInfo := <-deleteCommentLikeInfo
	if commentInfo == ERROR {
		log.Println("删除用户回复有误")
	}
	if likeInfo == ERROR {
		log.Println("删除用户like有误")
	}
	if err != nil {
		return ERROR
	} else {
		log.Println("用户删除" + strconv.Itoa(id))
		return SUCCESS
	}
}

func GetUserCommentImageUrl(id int) string {
	sql := "select comment_image_url from user_comment where id=?"
	var url string
	err := DB.QueryRow(sql, id).Scan(&url)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	return url
}

func UpdateUserCommentImage(id int, imgString string) string {
	img, err := base64.StdEncoding.DecodeString(imgString)
	if err != nil {
		log.Println("图片有误，base转换失败")
		return ERROR
	}
	isSave, commentImageUrl := SaveFindImage(img)
	if !isSave {
		log.Println("保存图片失败")
		return ERROR
	}
	sql := "update user_comment set comment_image_url=? where id=?"
	_, err = DB.Exec(sql, commentImageUrl, id)
	if err != nil {
		log.Println("更新帖子：" + strconv.Itoa(id) + " 失败")
		return ERROR
	} else {
		log.Println("更新帖子：" + strconv.Itoa(id) + " 成功")
		return SUCCESS
	}
}

func UpdateUserCommentInformaiton(id int,
	commentTitle string,
	comemntContent string,
	commentImage string) string {
	sql := "update user_comment set comment_title=?,comment_content=?,comment_image_url=? where id=?"
	imgByte, err := base64.StdEncoding.DecodeString(commentImage)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	isSave, commentImageUrl := SaveFindImage(imgByte)
	if !isSave {
		log.Println("保存图片错误")
		return ERROR
	}
	result, err := DB.Exec(sql, commentTitle, comemntContent, commentImageUrl, id)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affect, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	if affect > 0 {
		return commentImageUrl
	}
	return ERROR
}

func GetUserCommentIdByUser(userID int) string {
	sql := "SELECT COUNT(id) from user_comment where user_id=? ORDER BY id DESC"
	var count int
	err := DB.QueryRow(sql, userID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	sql = "SELECT id from user_comment where user_id=? ORDER BY id DESC"
	rows, err := DB.Query(sql, userID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]int, count)
	for i := 0; rows.Next(); i++ {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		resultList[i] = id
	}
	resultjson, err := json.Marshal(resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(resultjson)
	return jsonString
}

func GetUserCommentInformationByUser(userID int, start int) string {
	sql := "SELECT user_comment.id,user_id,user_name,comment_time,comment_title,comment_content,comment_image_url from user_comment,user_info where user_id in (select focus_fansID from my_focus where focus_userID=?) and user_id=user_info.ID order by id DESC LIMIT ?,20"
	rows, err := DB.Query(sql, userID, start)
	defer rows.Close()
	if err != nil {
		log.Panicln(err.Error())
		return ERROR
	}
	resultList := make([]UserCommentData, 20)
	var count = 0
	for i := 0; rows.Next(); i++ {
		var data UserCommentData
		err := rows.Scan(&data.ID, &data.UserID, &data.UserName, &data.CommentTime, &data.CommentTitle, &data.ComemntContent, &data.ImageURL)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		data.IsLike = GetUserIsLike(userID, data.ID)
		data.IsFollow = IsUserFollow(userID, data.UserID)
		sql = "SELECT COUNT(*) from user_comment_like where commentID=?"
		err = DB.QueryRow(sql, data.ID).Scan(&data.LikeNum)
		if err != nil {
			log.Println(err.Error())
			data.LikeNum = 0
		}
		sql = "SELECT COUNT(*) from user_comment_reply where comment_ID=?"
		err = DB.QueryRow(sql, data.ID).Scan(&data.ReplyNum)
		if err != nil {
			log.Println(err.Error())
			data.ReplyNum = 0
		}
		resultList[i] = data
		count = i
	}
	jsonResult, err := json.Marshal(resultList[0 : count+1])
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	return jsonString
}

func GetUserCommentInformaitonByOwn(userID int, start int) string {
	sql := "SELECT user_comment.id,user_id,user_name,comment_time,comment_title,comment_content,comment_image_url " +
		"from user_comment,user_info where user_id=? and user_id=user_info.id order by id DESC LIMIT ?,20"
	rows, err := DB.Query(sql, userID, start)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]UserCommentData, 20)
	var count = 0
	for i := 0; rows.Next(); i++ {
		var data UserCommentData
		err := rows.Scan(&data.ID, &data.UserID, &data.UserName, &data.CommentTime, &data.CommentTitle, &data.ComemntContent, &data.ImageURL)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		data.IsLike = GetUserIsLike(userID, data.ID)
		data.IsFollow = IsUserFollow(userID, data.UserID)
		sql = "SELECT COUNT(*) from user_comment_like where commentID=?"
		err = DB.QueryRow(sql, data.ID).Scan(&data.LikeNum)
		if err != nil {
			log.Println(err.Error())
			data.LikeNum = 0
		}
		sql = "SELECT COUNT(*) from user_comment_reply where comment_ID=?"
		err = DB.QueryRow(sql, data.ID).Scan(&data.ReplyNum)
		if err != nil {
			log.Println(err.Error())
			data.ReplyNum = 0
		}
		resultList[i] = data
		count = i
	}
	jsonResult, err := json.Marshal(resultList[0 : count+1])
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	return jsonString
}

func getAllUserCommentInfoClass(userID int, commentID int) (UserCommentData, error) {
	sql := "SELECT user_comment.id,user_id,user_name,comment_time,comment_title,comment_content,comment_image_url from user_comment,user_info where user_comment.id=? and user_id=user_info.id"
	var data UserCommentData
	err := DB.QueryRow(sql, commentID).Scan(&data.ID, &data.UserID, &data.UserName, &data.CommentTime, &data.CommentTitle, &data.ComemntContent, &data.ImageURL)
	if err != nil {
		log.Println(err.Error())
		return data, err
	}
	data.ReplyNum = GetUserCommentCount(commentID, normalReply)
	data.IsLike = GetUserIsLike(userID, commentID)
	data.LikeNum = GetCommentLikeNumber(commentID)
	data.IsFollow = IsUserFollow(userID, data.UserID)
	data.IsCollect = CheckIsCollection(userID, TYPE_FIND, commentID)
	return data, nil
}

func GetAllUserCommentInfoByID(user int, commentID int) string {
	sql := "SELECT user_comment.id,user_id,user_name,comment_time,comment_title,comment_content,comment_image_url from user_comment,user_info where user_comment.id=? and user_id=user_info.id"
	var data UserCommentData
	err := DB.QueryRow(sql, commentID).Scan(&data.ID, &data.UserID, &data.UserName, &data.CommentTime, &data.CommentTitle, &data.ComemntContent, &data.ImageURL)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	data.ReplyNum = GetUserCommentCount(commentID, normalReply)
	data.IsLike = GetUserIsLike(data.UserID, commentID)
	data.LikeNum = GetCommentLikeNumber(commentID)
	data.IsFollow = IsUserFollow(user, data.UserID)
	jsonResult, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	return jsonString
}

func GetCommentLikeNumber(commentID int) int {
	sql := "SELECT COUNT(*) from user_comment_like where commentID=?"
	var count int
	err := DB.QueryRow(sql, commentID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return count
}

func GetUserCommentCount(commentID int, replyType int) int {
	if replyType == miniReply {
		return 3
	}
	sql := "SELECT count(*) from user_comment_reply where comment_ID=?"
	var count int
	err := DB.QueryRow(sql, commentID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return count
}

func GetUserCommentReply(commentID int, replyType int) string {
	count := GetUserCommentCount(commentID, replyType)
	if count == 0 {
		return ERROR
	}
	sql := ""
	if replyType == normalReply {
		sql = "SELECT reply_id,reply_time,comment_ID,user_id,user_name,reply_content from user_comment_reply,user_info where comment_ID=? and user_info.id=user_id"
	} else {
		sql = "SELECT reply_id,reply_time,comment_ID,user_id,user_name,reply_content from user_comment_reply,user_info where comment_ID=? and user_info.id=user_id order by reply_id desc limit 0,3"
	}
	rows, err := DB.Query(sql, commentID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]ReplyInformation, count)
	for i := 0; rows.Next(); i++ {
		var data ReplyInformation
		err := rows.Scan(&data.ID, &data.ReplyTime, &data.CommentID, &data.UserID, &data.UserName, &data.ReplyContent)
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
	return jsonString
}

func SetUserLike(userID int, commentID int) string {
	sql := "insert into user_comment_like(userID,commentID) VALUES(?,?) "
	_, err := DB.Exec(sql, userID, commentID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	return SUCCESS
}

func CancelUserLike(userID int, commentID int) string {
	sql := "delete from user_comment_like where userID=? and commentID=?"
	result, err := DB.Exec(sql, userID, commentID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	rowsAffected, err := result.RowsAffected()
	if err == nil && rowsAffected >= 1 {
		return SUCCESS
	}
	return ERROR
}

func AddUserCommentReply(userID int, commentID int, replyContent string) string {
	sql := "insert into user_comment_reply(user_id,comment_id,reply_time,reply_content) values(?,?,?,?)"
	thisTime := time.Now().Format("2006-01-02 03:04:05")
	result, err := DB.Exec(sql, userID, commentID, thisTime, replyContent)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affectCount, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultString := strconv.FormatInt(affectCount, 10)
	return resultString
}

func UpdateUserCommentReply(replyID int, replyContent string) string {
	sql := "UPDATE user_comment_reply SET reply_content=? where reply_id=?"
	result, err := DB.Exec(sql, replyContent, replyID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affectCount, err := result.RowsAffected()
	if affectCount >= 0 && err == nil {
		return SUCCESS
	}
	return ERROR
}

func DeleteUserCommentReply(replyID int) string {
	sql := "delete from user_comment_reply where reply_id=?"
	result, err := DB.Exec(sql, replyID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affectCount, err := result.RowsAffected()
	if affectCount >= 0 && err == nil {
		return SUCCESS
	}
	return ERROR
}

func GetUserLikeComment(userID int) string {
	sql := "select COUNT(id) from user_comment_like where userID=?"
	var count int
	err := DB.QueryRow(sql, userID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	sql = "select commentID from user_comment_like where userID=?"
	rows, err := DB.Query(sql, userID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	defer rows.Close()
	var resultList = make([]UserCommentData, count)
	for i := 0; rows.Next(); i++ {
		var commentID int
		err := rows.Scan(&commentID)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		data, err := getAllUserCommentInfoClass(userID, commentID)
		if err != nil {
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
	log.Println("用户：" + strconv.Itoa(userID) + " 获取我的赞")
	return jsonString
}
