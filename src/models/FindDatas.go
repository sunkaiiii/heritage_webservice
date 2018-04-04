package models

type FindActivityData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
type UserCommentData struct {
	ID             int    `json:"id"`
	CommentTime    string `json:"commentTime"`
	UserName       string `json:"userName"`
	CommentTitle   string `json:"commentTitle"`
	ComemntContent string `json:"commentContent"`
	UserID         int    `json:"userID"`
	IsLike         string `json:"isLike"`
	IsFollow       string `json:"isFollow"`
	IsCollect      string `json:"isCollect"`
	LikeNum        int    `json:"likeNum"`
	ReplyNum       int    `json:"replyNum"`
	ImageURL       string `json:"imageUrl"`
}

type ReplyInformation struct {
	ID           int    `json:"id"`
	ReplyTime    string `json:"replyTime"`
	CommentID    int    `json:"commentID"`
	UserID       int    `json:"userID"`
	UserName     string `json:"userName"`
	ReplyContent string `json:"replyContent"`
}
