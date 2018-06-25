package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
)

const push_user_id = "push_user_id:"

var socket_map map[int]net.Conn

func AddPushMessage(userID int, userName string, replyCommentID int, replyContent string, replyToUserID int, replyToUserName string, replyTime string, originaReplyContent string) string {
	sql := "insert into user_comment_push_reply(userID,userName,replyCommentID,replyContent,replyToUserID,replyToUserName,replyTime,originalReplyContent) values(?,?,?,?,?,?,?,?)"
	result, err := DB.Exec(sql, userID, userName, replyCommentID, replyContent, replyToUserID, replyToUserName, replyTime, originaReplyContent)
	rowID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	go sendPush(rowID, userID, userName, replyCommentID, replyContent, replyToUserID, replyToUserName, replyTime, originaReplyContent)
	return SUCCESS
}

func sendPush(id int64, userID int, userName string, replyCommentID int, replyContent string, replyToUserID int, replyToUserName string, replyTime string, originaReplyContent string) {
	value, err := GetRedisKey(push_user_id + strconv.Itoa(replyToUserID))
	if err != nil {
		log.Println(err.Error())
		return
	}
	valueNum, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	socketConnect, ok := socket_map[valueNum]
	if !ok {
		return
	}
	var data PushMessageData
	data.ID = int(id)
	data.UserName = userName
	data.ReplyCommentID = replyCommentID
	data.ReplyContent = replyContent
	data.ReplyTime = replyTime
	data.OriginalReplyContent = originaReplyContent
	jsonResult, err := json.Marshal(data)
	if err == nil {
		socketConnect.Write([]byte(string(jsonResult)))
		log.Println("发送推送向" + strconv.Itoa(replyToUserID) + "内容：" + string(jsonResult))
		updateIsPushStateBySingle(id)
	}
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

func SendSinglePushMessageInfo(userID int) string {
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
	sql := "select id,userID,userName,replyCommentID,replyContent,replyTime,originalReplyContent from user_comment_push_reply where replyToUserID=? order by id DESC"
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

func StartPushService() {
	socket_count := 0
	socket_map = map[int]net.Conn{}
	listen, err := net.Listen("tcp", ":8088")
	defer listen.Close()
	if err != nil {
		log.Println("listen error " + err.Error())
		return
	}
	log.Println("Socket等待连接....")
	for {
		c, err := listen.Accept()
		if err != nil {
			log.Println("accept error " + err.Error())
			return
		}
		socket_count++
		socket_map[socket_count] = c
		go handleConnect(c, socket_count)
	}
}

func handleConnect(c net.Conn, socketID int) {
	defer c.Close()
	log.Println(strconv.Itoa(socketID))
	c.Write([]byte(fmt.Sprintf("%d\n", socketID)))
	buf := make([]byte, 50)
	n, err := c.Read(buf)
	if err != nil {
		log.Println(err.Error())
		return
	}
	userID := string(buf[0:n])
	log.Println("用户连接:" + userID)
	defer log.Println("socket:" + strconv.Itoa(socketID) + " close")
	SetRedisKey(push_user_id+userID, strconv.Itoa(socketID))
	defer DeleteRedisKey(push_user_id + userID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(buf[0:n]))
	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}
		log.Println(string(buf[0:n]))
	}
}

func updateIsPushStateBySingle(id int64) {
	sql := "update user_comment_push_reply set isPush=1 where id=?"
	DB.Exec(sql, id)
}

func updateIsPushState(userID int) {
	sql := "update user_comment_push_reply set isPush=1 where replyToUserID=?"
	DB.Exec(sql, userID)
}
