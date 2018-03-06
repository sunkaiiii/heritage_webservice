package models

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"time"
)

func GetFollowNumber(userID int) int {
	sql := "select count(id) from my_focus where focus_userID=?"
	var count int
	err := DB.QueryRow(sql, userID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return count
}

func GetFansNumber(userID int) int {
	sql := "select count(id) from my_focus where focus_fansID=?"
	var count int
	err := DB.QueryRow(sql, userID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return count
}

func GetUserPermission(userID int) int {
	sql := "select USER_PERMISSION from user_info where id=?"
	var permission int
	err := DB.QueryRow(sql, userID).Scan(&permission)
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return permission
}

func GetUserFocusAndFansViewPermission(userID int) int {
	sql := "select USER_FOCUS_AND_FANS_VIEW_PERMISSION from user_info where id=?"
	var permission int
	err := DB.QueryRow(sql, userID).Scan(&permission)
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return permission
}

func SetUserPermission(userID int, permission int) string {
	sql := "UPDATE user_info set USER_PERMISSION=? where id=?"
	result, err := DB.Exec(sql, permission, userID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affectedRow, err := result.RowsAffected()
	if err == nil && affectedRow > 0 {
		return SUCCESS
	}
	log.Println(err.Error())
	return ERROR
}

func SetUserFocusAndFansViewPermission(userID int, permission int) string {
	sql := "UPDATE user_info set USER_FOCUS_AND_FANS_VIEW_PERMISSION=? where id=?"
	result, err := DB.Exec(sql, permission, userID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affectedRow, err := result.RowsAffected()
	if err == nil && affectedRow > 0 {
		return SUCCESS
	}
	log.Println(err.Error())
	return ERROR
}

func GetUserAllInfo(userID int) string {
	var data UserInfo
	data.ID = userID
	data.UserName = getUserNameByUserID(userID)
	data.FocusNumber = GetFollowNumber(userID)
	data.FansNumber = GetFansNumber(userID)
	data.Permission = GetUserPermission(userID)
	data.FocusAndFansPermission = GetUserFocusAndFansViewPermission(userID)
	jsonResult, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	jsonString := string(jsonResult)
	return jsonString
}

func GetFollowInformation(userID int) string {
	count := GetFollowNumber(userID)
	if count == 0 {
		return ERROR
	}
	sql := "select focus_fansID,USER_NAME,focus_userID from my_focus,user_info where my_focus.focus_userID=? and my_focus.focus_fansID=user_info.ID"
	rows, err := DB.Query(sql, userID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]FollowInformation, count)
	for i := 0; rows.Next(); i++ {
		var data FollowInformation
		err := rows.Scan(&data.FocusFansID, &data.UserName, &data.FocusFocusID)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		data.Checked = (SUCCESS == IsUserFollow(data.FocusFocusID, data.FocusFansID))
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

func GetFansInformation(userID int) string {
	count := GetFansNumber(userID)
	if count == 0 {
		return ERROR
	}
	sql := "select focus_userID,USER_NAME,focus_fansID from my_focus,user_info where my_focus.focus_fansID=? and my_focus.focus_userID=user_info.ID"
	rows, err := DB.Query(sql, userID)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]FollowInformation, count)
	for i := 0; rows.Next(); i++ {
		var data FollowInformation
		err := rows.Scan(&data.FocusFansID, &data.UserName, &data.FocusFocusID)
		if err != nil {
			log.Println(err.Error())
			return ERROR
		}
		data.Checked = (SUCCESS == IsUserFollow(data.FocusFocusID, data.FocusFansID))
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

func AddFocus(userID int, focusID int) string {
	sql := "insert into my_focus(focus_userID,focus_fansID) VALUES (?,?)"
	result, err := DB.Exec(sql, userID, focusID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	afftectNum, err := result.RowsAffected()
	if err == nil && afftectNum > 0 {
		return SUCCESS
	}
	log.Println(err.Error())
	return ERROR
}

func CancelFocus(userID int, focusID int) string {
	sql := "delete from my_focus where focus_userID=? and focus_fansID=?"
	result, err := DB.Exec(sql, userID, focusID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	afftectNum, err := result.RowsAffected()
	if err == nil && afftectNum > 0 {
		return SUCCESS
	}
	log.Println(err.Error())
	return ERROR
}

func CheckFollowEachother(userID int, focusID int) string {
	sql := "SELECT count(id) from my_focus where focus_userID=? and focus_fansID=?"
	var count int
	err := DB.QueryRow(sql, focusID, userID).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	if count > 0 {
		return SUCCESS
	}
	return ERROR
}

func GetSearchUserInfo(name string) string {
	searchName := "%" + name + "%"
	sql := "select COUNT(id) from user_info where user_name like ?"
	var count int
	err := DB.QueryRow(sql, searchName).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	if count == 0 {
		return ERROR
	}
	sql = "select id,user_name from user_info where user_name like ?"
	rows, err := DB.Query(sql, searchName)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	resultList := make([]SearchInfo, count)
	for i := 0; rows.Next(); i++ {
		var data SearchInfo
		err := rows.Scan(&data.ID, &data.UserName)
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

func GetUserImage(userID int) string {
	sql := "select USER_IS_HAD_IMAGE,user_image_url from user_info where id=?"
	var hadImage int
	var imageURL string
	err := DB.QueryRow(sql, userID).Scan(&hadImage, &imageURL)
	if err != nil {
		log.Println("load user image error:" + err.Error())
		return ERROR
	}
	if hadImage == 0 {
		return ERROR
	}
	return imageURL
}

func UpdateUserImage(userID int, imageString string) string {
	sql := "UPDATE USER_INFO set user_image=?,user_image_url=?,USER_IS_HAD_IMAGE=?,USER_UPDATE_TIME=? where id=?"
	imgByte, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		log.Println("用户图片错误:" + err.Error())
		return ERROR
	}
	isSave, sqlUrl := saveUserImage(imgByte)
	if !isSave {
		log.Println("保存图片失败")
		return ERROR
	}
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	hadImage := 1
	result, err := DB.Exec(sql, imgByte, sqlUrl, hadImage, updateTime, userID)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	affected, err := result.RowsAffected()
	if affected > 0 {
		return SUCCESS
	}
	log.Println(err.Error())
	return ERROR
}

func getUserIdFromCommentID(commentID int) int {
	sql := "select user_id from user_comment where id=?"
	var id int
	err := DB.QueryRow(sql, commentID).Scan(&id)
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return id
}

func getUserNameByUserID(userID int) string {
	sql := "select USER_NAME from user_info where id=?"
	var userName string
	err := DB.QueryRow(sql, userID).Scan(&userName)
	if err != nil {
		log.Println(err.Error())
		return ERROR
	}
	return userName
}
