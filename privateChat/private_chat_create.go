package privateChat

import (
	"moments/db"
	"moments/ent"
)

func CreatePrivateChat(connection *db.DatabaseConnection, firstUser, secondUser *ent.User) (*ent.PrivateChat, error) {
	return connection.Client.PrivateChat.Create().SetFirstUser(firstUser).SetSecondUser(secondUser).Save(*connection.Ctx)
}
