package api

import "github.com/gin-gonic/gin"

func userGroup(group *gin.RouterGroup) {
	auth := group.Group("/users")
	auth.DELETE("/", deleteUser())
}

func deleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
