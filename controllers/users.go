package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moments/api/auth"
	"moments/db"
	"moments/ent"
	"moments/user"
	"moments/utils"
	"net/http"
)

func checkUserCredentials(u *ent.User) (gin.H, int) {
	if len(u.Username) == 0 {
		return gin.H{
			"message": "json body does not contain any username",
		}, http.StatusBadRequest
	}
	if len(u.Password) == 0 {
		return gin.H{
			"message": "json body does not contain any password",
		}, http.StatusBadRequest
	}
	return nil, -1
}

// Register opens a db handle and create a new user
// returns gin.H and int as response body and status code
func Register(u *ent.User) (gin.H, int) {
	resp, statusCode := checkUserCredentials(u)
	if statusCode == -1 {
		return resp, statusCode
	}
	conn, cancel := db.New()
	defer cancel()
	defer conn.Client.Close()

	jwtService := auth.New() // create new instance to work with jwt service

	createdUser, err := user.CreateUser(conn, u) // create user in database
	if err != nil {
		return checkErrors(err)
	}

	token := jwtService.GenerateToken(createdUser.Email)
	return gin.H{
		"message": "user created successfully",
		"user":    createdUser,
		"token":   token,
	}, http.StatusCreated
}

//Login gets user information and return response and status code
func Login(u *ent.User) (gin.H, int) {
	resp, statusCode := checkUserCredentials(u)
	if statusCode == -1 {
		return resp, statusCode
	}
	conn, cancel := db.New()
	defer conn.Client.Close()
	defer cancel()
	fetchedUser, err := user.GetUserByUsername(conn, u.Username)
	if err != nil {
		return checkErrors(err)
	}
	if utils.CheckPasswordHash(u.Password, fetchedUser.Password) {
		jwtService := auth.New()
		token := jwtService.GenerateToken(u.Username)
		return gin.H{
			"message": "user logged in successfully",
			"token":   token,
			"user":    fetchedUser,
		}, http.StatusOK
	}
	return gin.H{
		"message": "wrong password",
	}, http.StatusBadRequest
}
// Delete deletes a user, if user with username exists in database
func Delete(username string) (gin.H, int) {
	conn, cancel := db.New()
	defer conn.Client.Close()
	defer cancel()

	fetchedUser, err := user.GetUserByUsername(conn, username)
	if err != nil {
		return checkErrors(err)
	}

	if err = user.DeleteUser(conn, fetchedUser); err != nil {
		return checkErrors(err)
	} else {
		return gin.H{
			"message": fmt.Sprintf("user with username %s deleted successfully", username),
			"user":    fetchedUser,
		}, http.StatusOK
	}

}

func Update(username string, u *ent.User) (gin.H, int) {
	conn, cancel := db.New()
	defer conn.Client.Close()
	defer cancel()

	updatedUser, err:=user.UpdateUserWithUsername(conn,u,username)
	if err != nil {
		return checkErrors(err)
	}
	return gin.H{
		"message": "user updated",
		"user": updatedUser,
	}, http.StatusOK
}

func ChangePassword(username string, newPassword, currentPassword string) (gin.H, int) {
	conn, cancel := db.New()
	defer conn.Client.Close()
	defer cancel()

	if len(newPassword) == 0 || len(currentPassword) == 0 {
		return gin.H{
			"message": "current or new password are empty",
		}, http.StatusBadRequest
	}
	_, err := user.ChangePassword(conn, username, newPassword, currentPassword)
	if err != nil {
		return checkErrors(err)
	}
	return gin.H{
		"message": "password changed successfully",
	}, http.StatusOK
}
