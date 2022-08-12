package controllers

import (
	"github.com/gin-gonic/gin"
	"moments/api/auth"
	"moments/db"
	"moments/ent"
	"moments/user"
	"net/http"
)

// Register opens a db handle and create a new user
// returns gin.H and int as response body and status code
func Register(u *ent.User) (gin.H, int) {
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
