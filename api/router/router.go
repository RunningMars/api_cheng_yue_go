package router

import (
	"api_go/api/controller"
	"api_go/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	r.POST("/user/login", controller.Login)
	r.POST("/user/register", controller.Register)
	r.Any("/sms/validate/send", controller.SendOtp)

	r.Use(middleware.AuthMiddleware())
	r.Any("/member/test", controller.Test)

	r.Any("/member/list", controller.MemberList)
	r.Any("/member/detail", controller.MemberDetail)
	r.POST("/member/save", controller.MemberSave)
	r.Any("/member/thumbs_up/list", controller.ThumbsUpList)
	r.Any("/member/favorite/update", controller.UpdateFavorite)
	r.Any("/member/thumbs_up/update", controller.UpdateThumbsUp)

	r.Any("/message/chat/list", controller.ChatList)
	r.Any("/message/chat/message", controller.MessageList)
	r.Any("/message/chat/send", controller.SendMessage)
	r.Any("/message/chat/unread_count", controller.UnreadCount)
	r.Any("/message/chat/read_all", controller.ReadAll)

	r.POST("/uploadImage", controller.UploadImage)

	r.Any("/user/logout", controller.Logout)

	r.GET("/test", controller.Test)
	r.GET("/index", controller.Index)

	r.GET("/getTest", controller.GetUserTest)
	r.GET("/getUsers", controller.GetUsers)

	// api := r.Group("/api")
	// {
	// 	api.GET("/users", controller.Index)
	// 	api.POST("/users", controller.CreateUser)
	// }
}
