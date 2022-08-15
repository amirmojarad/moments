package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moments/db"
	"moments/room"
	"moments/user"
	"net/http"
	"strings"
)

func CreatePrivateChat(firstUsername string, usernames []string) (gin.H, int) {
	if len(firstUsername) == 0 || len(usernames) == 0 {
		return gin.H{
			"message": "first username or second username is empty",
		}, http.StatusBadRequest
	}

	conn, cancel := db.New()
	defer conn.Client.Close()
	defer cancel()
	firstUser, err := user.GetUserByUsername(conn, firstUsername)
	if err != nil {
		return checkErrors(err)
	}
	secondUser, err := user.GetUserByUsername(conn, usernames[0])
	if err != nil {
		return checkErrors(err)
	}
	createdRoom, err := room.CreatePrivateRoom(conn, firstUser, secondUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return gin.H{
				"message": "duplicate private chat with given users",
				"error":   err.Error(),
			}, http.StatusConflict
		}
		return checkErrors(err)
	}
	users, err := room.GetUsersOfPrivateChat(conn, createdRoom.ID)
	if err != nil {
		return checkErrors(err)
	}
	return gin.H{
		"created_room": createdRoom,
		"message":      "private chat created successfully.",
		"users":        users,
	}, http.StatusCreated
}

func GetAllUserPrivateChats(username string) (gin.H, int) {
	if len(username) == 0 {
		return gin.H{
			"message": "request does not contain any token.",
		}, http.StatusBadRequest
	}

	conn, cancel := db.New()
	defer conn.Client.Close()
	defer cancel()

	userByUsername, err := user.GetUserByUsername(conn, username)
	if err != nil {
		return checkErrors(err)
	}
	privateChats, err := room.GetAllUserPrivateRooms(conn, userByUsername)
	if err != nil {
		return checkErrors(err)
	}
	return gin.H{
		"private_chats": privateChats,
		"message":       fmt.Sprintf("all private chats of user %s fetched successfully.", username),
	}, http.StatusOK
}
