package api

import "moments/ent"

type CreateRoomSchema struct {
	Users []string `json:"users"`
	Type  string   `json:"type"`
}

type AuthResponseSchema struct {
	Token   string    `json:"token"`
	Message string    `json:"message"`
	User    *ent.User `json:"user"`
}
