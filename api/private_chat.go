package api

import (
	"github.com/gin-gonic/gin"
	"moments/api/middlewares"
	"moments/controllers"
)

func privateChatGroup(group *gin.RouterGroup) {
	privateChat := group.Group("/private_chat", middlewares.CheckAuth())
	privateChat.POST("/", createPrivateChat())
	privateChat.GET("/all", getAllUserPrivateChats())
	privateChat.GET("/:id", getPrivateChat())
	privateChat.DELETE("/:id", deletePrivateChat())
	privateChat.POST("/:id/messages", addMessageToPrivateChat())
	privateChat.DELETE("/:id/messages", deleteMessageFromPrivateChat())
}

func createPrivateChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsernameFromCtx(ctx)
		var schema CreateRoomSchema
		ctx.BindJSON(&schema)
		response, statusCode := controllers.CreatePrivateChat(username, schema.Usernames)
		ctx.IndentedJSON(statusCode, response)
	}
}
func getAllUserPrivateChats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func getPrivateChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func deletePrivateChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func addMessageToPrivateChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func deleteMessageFromPrivateChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
