package messages

import (
	"moments/db"
	"moments/ent"
)

func GetAllUserMessages(conn *db.DatabaseConnection, user *ent.User) ([]*ent.Message, error) {
	return user.QueryMessages().All(*conn.Ctx)
}

func GetAllRoomMessages(conn *db.DatabaseConnection, room *ent.Room) ([]*ent.Message, error) {
	return room.QueryMessages().All(*conn.Ctx)
}

func GetMessageRoomOwner(conn *db.DatabaseConnection, message *ent.Message) (*ent.Room, error) {
	return message.QueryOwner().First(*conn.Ctx)
}

func GetMessageSenderUser(conn *db.DatabaseConnection, message *ent.Message) (*ent.User, error) {
	return message.QuerySender().First(*conn.Ctx)
}
