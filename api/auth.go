package api

import (
	"github.com/gin-gonic/gin"
	"moments/controllers"
	"moments/ent"
)

func authGroup(group *gin.RouterGroup) {
	auth := group.Group("/auth")
	auth.POST("/register", register())
	auth.POST("/login", login())
}

func register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user ent.User
		ctx.BindJSON(&user)
		response, status := controllers.Register(&user)
		ctx.IndentedJSON(status, response)
	}
}

func login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user ent.User
		ctx.BindJSON(&user)
		response, status := controllers.Login(&user)
		ctx.IndentedJSON(status, response)
	}
}
