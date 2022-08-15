package controllers

import (
	"github.com/gin-gonic/gin"
	"moments/db"
	"moments/room"
	"moments/user"
	"net/http"
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
		return checkErrors(err)
	}
	users, err := room.GetPrivateUserParticipants(conn, createdRoom.ID)
	if err != nil {
		return checkErrors(err)
	}
	return gin.H{
		"created_room": createdRoom,
		"message":      "private chat created successfully.",
		"users":        users,
	}, http.StatusCreated
}
