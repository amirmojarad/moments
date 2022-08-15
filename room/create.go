package room

import (
	"errors"
	"moments/db"
	"moments/ent"
	"moments/ent/room"
)

func CreatePrivateRoom(connection *db.DatabaseConnection, users ...*ent.User) (*ent.Room, error) {
	if ok := ExistRoomWithUsers(connection, users[0], users[1]); ok {
		return nil, errors.New("duplicate private room with given users")
	}
	return connection.Client.Room.Create().
		SetTitle("").
		SetType(room.TypePrivate).
		AddUsers(users...).Save(*connection.Ctx)
}

func CreatePublicRoom(connection *db.DatabaseConnection, groupTitle string, users ...*ent.User) (*ent.Room, error) {
	return connection.Client.Room.Create().SetTitle(groupTitle).AddUsers(users...).
		Save(*connection.Ctx)
}
