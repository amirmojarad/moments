package api

import "moments/ent"

type CreateRoomSchema struct {
	Usernames []string `json:"usernames"`
}

type AuthResponseSchema struct {
	Token   string    `json:"token"`
	Message string    `json:"message"`
	User    *ent.User `json:"user"`
}

type CreateRoomResponseSchema struct {
	CreatedRoom ent.Room   `json:"created_room"`
	Message     string     `json:"message"`
	Users       []ent.User `json:"users"`
}
