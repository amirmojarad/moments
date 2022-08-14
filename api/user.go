package api

import (
	"github.com/gin-gonic/gin"
	"moments/api/middlewares"
	"moments/controllers"
	"moments/ent"
)

func userGroup(group *gin.RouterGroup) {
	user := group.Group("/users", middlewares.CheckAuth())
	user.POST("/change_password", changePassword())
	user.PUT("/", updateUser())
	user.PATCH("/", updateUser())
	user.DELETE("/", deleteUser())
}

func deleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsernameFromCtx(ctx)
		responseBody, statusCode := controllers.Delete(username)
		ctx.IndentedJSON(statusCode, responseBody)
	}
}

func updateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user ent.User
		ctx.BindJSON(&user)
		username := getUsernameFromCtx(ctx)
		responseBody, statusCode := controllers.Update(username, &user)
		ctx.IndentedJSON(statusCode, responseBody)
	}
}

func changePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type changePasswordRequestSchema struct {
			NewPassword     string `json:"new_password"`
			CurrentPassword string `json:"current_password"`
		}
		var schema changePasswordRequestSchema
		ctx.BindJSON(&schema)
		username := getUsernameFromCtx(ctx)
		responseBody, statusCode := controllers.ChangePassword(username, schema.NewPassword, schema.CurrentPassword)
		ctx.IndentedJSON(statusCode, responseBody)
	}

}
