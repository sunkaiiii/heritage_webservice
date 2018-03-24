package models

import (
	"database/sql"
	"encoding/json"
	"log"
)

func AddPushMessage(userID int, userName string, replyCommentID int, replyContent string, replyToUserID int, replyToUserName string, replyTime string, originaReplyContent string) string {
	sql := "insert into user_comment_push_reply(userID,userName,replyCommentID,replyContent,replyToUserID,replyToUserName,replyTime,originalReplyContent) values(?,?,?,?,?,?,?,?)"
	_, err := DB.Exec(sql, userID, userName, replyCommentID, replyContent, replyToUserID, replyToUserName, replyTime, originaReplyContent)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	return SUCCESS
}

func updateIsPushState(userID int) {
	sql := "update user_comment_push_reply set isPush=1 where replyToUserID=?"
	DB.Exec(sql, userID)
}

func getPushMessageInfo(rows *sql.Rows, count int) ([]PushMessageData, error) {
	resultList := make([]PushMessageData, count)
	i := 0
	for ; rows.Next(); i++ {
		var data PushMessageData
		err := rows.Scan(&data.ID, &data.UserID, &data.UserName, &data.ReplyCommentID, &data.ReplyContent, &data.ReplyTime, &data.OriginalReplyContent)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		resultList[i] = data
	}
	return resultList[0:i], nil
}

func SendPush(userID int) string {
	sql := "select id,userID,userName,replyCommentID,replyContent,replyTime,originalReplyContent from user_comment_push_reply where replyToUserID=? and isPush=0 LIMIT 0,3"
	rows, err := DB.Query(sql, userID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList, err := getPushMessageInfo(rows, 3)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	go updateIsPushState(userID)
	jsonResult, err := json.Marshal(resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	return string(jsonResult)
}

func GetPushMessageInfo(userID int) string {
	sql := "select id,userID,userName,replyCommentID,replyContent,replyTime,originalReplyContent from user_comment_push_reply where replyToUserID=?"
	rows, err := DB.Query(sql, userID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList, err := getPushMessageInfo(rows, 200)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonResult, err := json.Marshal(resultList)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	return jsonString
}
