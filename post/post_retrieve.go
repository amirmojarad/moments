package post

import (
	"moments/db"
	"moments/ent"
)

func GetPostOwner(connection *db.DatabaseConnection, post *ent.Post) (*ent.User, error) {
	return post.QueryOwner().Only(*connection.Ctx)
}

func GetPostByID(connection *db.DatabaseConnection, id int) (*ent.Post, error) {
	return connection.Client.Post.Get(*connection.Ctx, id)
}
