package messages

import (
	"moments/db"
	"moments/ent"
)

func Create(conn *db.DatabaseConnection, senderUser *ent.User, ownerRoom *ent.Room, text string) (*ent.Message, error) {
	return conn.Client.Message.Create().SetOwner(ownerRoom).SetSender(senderUser).SetText(text).Save(*conn.Ctx)
}
