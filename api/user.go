package api

import (
	"github.com/gin-gonic/gin"
	"moments/api/middlewares"
	"moments/controllers"
	"moments/ent"
	"net/http"
	"strconv"
)

func userGroup(group *gin.RouterGroup) {
	user := group.Group("/users", middlewares.CheckAuth())
	user.POST("/change_password", changePassword())
	user.PUT("/", updateUser())
	user.PATCH("/", updateUser())
	user.DELETE("/", deleteUser())

	// Follow System
	user.POST("/follow", follow())
	user.DELETE("/follow", unfollow())
	user.GET("/followers", getAllFollowers())
	user.GET("/following", getAllFollowing())
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

func follow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usernames []string
		ctx.BindJSON(&usernames)
		username := getUsernameFromCtx(ctx)
		response, statusCode := controllers.Follow(username, usernames...)
		ctx.IndentedJSON(statusCode, response)
	}

}

func unfollow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usernames []string
		ctx.BindJSON(&usernames)
		username := getUsernameFromCtx(ctx)
		response, statusCode := controllers.Unfollow(username, usernames...)
		ctx.IndentedJSON(statusCode, response)
	}
}

func getAllFollowers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsernameFromCtx(ctx)
		page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error while fetching page in query",
				"error":   err,
			})
		}
		response, statusCode := controllers.GetAllUserFollowers(username, page)
		ctx.IndentedJSON(statusCode, response)
	}
}

func getAllFollowing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsernameFromCtx(ctx)
		page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error while fetching page in query",
				"error":   err,
			})
		}
		response, statusCode := controllers.GetAllUserFollowing(username, page)
		ctx.IndentedJSON(statusCode, response)
	}
}
