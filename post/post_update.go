package post

import (
	"moments/db"
	"moments/ent"
)

func UpdateLinkUrl(connection *db.DatabaseConnection, post *ent.Post, linkUrl string) (*ent.Post, error) {
	return post.Update().SetLinkURL(linkUrl).Save(*connection.Ctx)
}
