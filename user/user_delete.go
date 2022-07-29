package user

import (
	"moments/db"
	"moments/ent"
)

func DeleteUser(connection *db.DatabaseConnection, user *ent.User) error {
	return connection.Client.User.DeleteOne(user).Exec(*connection.Ctx)
}
