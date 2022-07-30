package post

import (
	"moments/db"
	"moments/ent"
)

func DeletePost(connection *db.DatabaseConnection, user *ent.User, post *ent.Post) error {
	return connection.Client.Post.DeleteOne(post).Exec(*connection.Ctx)
}
