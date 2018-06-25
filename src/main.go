package main

import (
	"io"
	"os"

	"./controllers"
	"./models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := models.InitDB()
	if db == nil {
		return
	}
	defer db.Close()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f) //将log打印到文件
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	go models.StartPushService()
	router.Static("/img", "./img") //发布图片文件夹
	router.GET("/GetUserName", controllers.GetUserName)
	router.POST("/Sign_In", controllers.Sign_In)
	router.POST("/UserRegist", controllers.UserRegist)
	router.GET("/FindPassWordQuestion", controllers.FindPassWordQuestion)
	router.POST("/CheckQuestionAnswer", controllers.CheckQuestionAnswer)
	router.POST("/ChangePassword", controllers.ChangePassword)
	router.GET("/GetMainDivideActivityImageUrl", controllers.GetMainDivideActivityImageUrl)
	router.GET("/GetChannelInformation", controllers.GetChannelInformation)
	router.GET("/GetChannelFolkInformation", controllers.GetChannelFolkInformation)
	router.GET("/GetChannelFolkSingleInformation", controllers.GetChannelFolkSingleInformation)
	router.GET("/SearchChannelForkInfo", controllers.SearchChannelForkInfo)
	router.GET("/GetFindActivityID", controllers.GetFindActivityID)
	router.POST("/AddUserCommentInformation", controllers.AddUserCommentInformation)
	router.GET("/IsUserFollow", controllers.IsUserFollow)
	router.GET("/GetUserCommentInformation", controllers.GetUserCommentInformation)
	router.GET("/GetFolkNewsList", controllers.GetFolkNewsList)
	router.GET("/GetFolkNewsInformation", controllers.GetFolkNewsInformation)
	router.GET("/GetBottomNewsLiteInformation", controllers.GetBottomNewsLiteInformation)
	router.GET("/GetBottomNewsInformationByID", controllers.GetBottomNewsInformationByID)
	router.GET("/DeleteUserCommentByID", controllers.DeleteUserCommentByID)
	router.GET("/GetUserCommentImageUrl", controllers.GetUserCommentImageUrl)
	router.POST("/UpdateUserCommentImage", controllers.UpdateUserCommentImage)
	router.POST("/UpdateUserCommentInformaiton", controllers.UpdateUserCommentInformaiton)
	router.GET("/GetMainPageSlideNewsInformation", controllers.GetMainPageSlideNewsInformation)
	router.GET("/GetUserCommentIdByUser", controllers.GetUserCommentIdByUser)
	router.GET("/GetUserCommentInformationByUser", controllers.GetUserCommentInformationByUser)
	router.GET("/GetUserCommentInformaitonByOwn", controllers.GetUserCommentInformaitonByOwn)
	router.GET("/GetUserCommentInformationBySameLocation", controllers.GetUserCommentInformationBySameLocation)
	router.GET("/GetAllUserCommentInfoByID", controllers.GetAllUserCommentInfoByID)
	router.GET("/GetCommentLikeNumber", controllers.GetCommentLikeNumber)
	router.GET("/GetUserCommentCount", controllers.GetUserCommentCount)
	router.GET("/GetUserCommentReply", controllers.GetUserCommentReply)
	router.GET("/SetUserLike", controllers.SetUserLike)
	router.GET("/CancelUserLike", controllers.CancelUserLike)
	router.POST("/AddUserCommentReply", controllers.AddUserCommentReply)
	router.POST("/UpdateUserCommentReply", controllers.UpdateUserCommentReply)
	router.GET("/DeleteUserCommentReply", controllers.DeleteUserCommentReply)
	router.GET("/GetFollowNumber", controllers.GetFollowNumber)
	router.GET("/GetFansNumber", controllers.GetFansNumber)
	router.GET("/GetUserPermission", controllers.GetUserPermission)
	router.GET("/GetUserFocusAndFansViewPermission", controllers.GetUserFocusAndFansViewPermission)
	router.GET("/SetUserPermission", controllers.SetUserPermission)
	router.GET("/SetUserFocusAndFansViewPermission", controllers.SetUserFocusAndFansViewPermission)
	router.GET("/GetUserAllInfo", controllers.GetUserAllInfo)
	router.GET("/GetFollowInformation", controllers.GetFollowInformation)
	router.GET("/GetFansInformation", controllers.GetFansInformation)
	router.GET("/AddFocus", controllers.AddFocus)
	router.GET("/CancelFocus", controllers.CancelFocus)
	router.GET("/CheckFollowEachother", controllers.CheckFollowEachother)
	router.GET("/GetSearchUserInfo", controllers.GetSearchUserInfo)
	router.GET("/GetUserImage", controllers.GetUserImage)
	router.POST("/UpdateUserImage", controllers.UpdateUserImage)
	router.GET("/AddUserCollection", controllers.AddUserCollection)
	router.GET("/GetUserCollection", controllers.GetUserCollection)
	router.GET("/GetUserLikeComment", controllers.GetUserLikeComment)
	router.GET("/CancelUserCollect", controllers.CancelUserCollect)
	router.GET("/CheckIsCollection", controllers.CheckIsCollection)
	router.POST("/AddPushMessage", controllers.AddPushMessage)
	router.GET("/GetPushMessage", controllers.GetPushMessage)
	router.GET("/GetAllPushReplyInfo", controllers.GetAllPushReplyInfo)
	router.GET("/SearchBottomNewsInformation", controllers.SearchBottomNewsInformation)
	router.GET("/SearchFolkNewsInformaiton", controllers.SearchFolkNewsInformaiton)
	router.GET("/SearchUserCommentInfo", controllers.SearchUserCommentInfo)
	go router.RunTLS(":8081", "./models/1_sunkai.xyz_bundle.crt", "./models/2_sunkai.xyz.key")
	router.Run()
}
