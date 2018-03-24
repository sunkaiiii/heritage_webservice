package models

type PushReplyData struct {
	ID                   int    `json:"id"`
	UserID               int    `json:"userID"`
	UserName             string `json:"userName"`
	ReplyCommentID       int    `json:"replyComemntID"`
	ReplyContent         string `json:"replyContent"`
	ReplyToUserID        int    `json:"replyToUserID"`
	ReplyToUserName      string `json:"replyToUserName"`
	ReplyTime            string `json:"replyTime"`
	IsPush               int    `json:"isPush"`
	OriginalReplyContent string `json:"originalReplyContent"`
}

type PushMessageData struct {
	ID                   int    `json:"id"`
	UserID               int    `json:"userID"`
	UserName             string `json"userName"`
	ReplyCommentID       int    `json:"replyComemntID"`
	ReplyContent         string `json:"replyContent"`
	ReplyTime            string `json:"replyTime"`
	OriginalReplyContent string `json:"originalReplyContent"`
}
