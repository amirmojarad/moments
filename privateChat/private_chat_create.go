package privateChat

import (
	"moments/db"
	"moments/ent"
)

func CreatePrivateChat(connection *db.DatabaseConnection, sender *ent.User, receiver *ent.User) (*ent.PrivateChat, error) {
	return connection.Client.PrivateChat.Create().SetReceiver(receiver).SetSender(sender).Save(*connection.Ctx)
}
