package privateChat

import (
	"moments/db"
	"moments/ent"
)

func AddMessageToPrivateChat(connection *db.DatabaseConnection,
	pvChat *ent.PrivateChat,
	message *ent.Message) (*ent.PrivateChat, error) {
	return pvChat.Update().AddChats(message).Save(*connection.Ctx)
}
