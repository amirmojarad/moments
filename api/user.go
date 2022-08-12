package api

import (
	"github.com/gin-gonic/gin"
	"moments/api/middlewares"
	"moments/controllers"
)

func userGroup(group *gin.RouterGroup) {
	auth := group.Group("/users", middlewares.CheckAuth())
	auth.DELETE("/", deleteUser())
}

func deleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsernameFromCtx(ctx)
		responseBody, statusCode := controllers.Delete(username)
		ctx.IndentedJSON(statusCode, responseBody)
	}
}
