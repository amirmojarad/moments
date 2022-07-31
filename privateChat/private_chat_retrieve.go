package privateChat

import (
	"moments/db"
	"moments/ent"
)

func GetPrivateChatByID(connection *db.DatabaseConnection, pvChatID int) (*ent.PrivateChat, error) {
	return connection.Client.PrivateChat.Get(*connection.Ctx, pvChatID)
}

func GetPrivateChatFirstUser(connection *db.DatabaseConnection, pvChat *ent.PrivateChat) (*ent.User, error) {
	return pvChat.QueryFirstUser().Only(*connection.Ctx)
}

func GetPrivateChatSecondUser(connection *db.DatabaseConnection, pvChat *ent.PrivateChat) (*ent.User, error) {
	return pvChat.QuerySecondUser().Only(*connection.Ctx)
}
