package room

import (
	"moments/db"
	"moments/ent"
)

func AddNewUsersToPublicRoom(conn *db.DatabaseConnection, publicRoom *ent.Room, users ...*ent.User) (*ent.Room, error) {
	return publicRoom.Update().AddUsers(users...).Save(*conn.Ctx)
}
