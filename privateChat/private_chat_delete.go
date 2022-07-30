package privateChat

import (
	"moments/db"
	"moments/ent"
)

func DeletePrivateChatByID(connection *db.DatabaseConnection, id int) error {
	return connection.Client.PrivateChat.DeleteOneID(id).Exec(*connection.Ctx)
}

func DeletePrivateChat(connection *db.DatabaseConnection, pvChat *ent.PrivateChat) error {
	return connection.Client.PrivateChat.DeleteOne(pvChat).Exec(*connection.Ctx)
}
