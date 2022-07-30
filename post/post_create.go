package post

import (
	"moments/db"
	"moments/ent"
)

func CreatePost(connection *db.DatabaseConnection, user *ent.User, post *ent.Post) (*ent.Post, error) {
	return connection.Client.Post.Create().
		SetPlainText(post.PlainText).
		SetOwner(user).
		Save(*connection.Ctx)
}
