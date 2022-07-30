package user

import (
	"moments/db"
	"moments/ent"
	"time"
)

func AddFullName(connection *db.DatabaseConnection, user *ent.User, fullName string) (*ent.User, error) {
	return user.Update().SetFullName(fullName).SetUpdatedDate(time.Now()).Save(*connection.Ctx)
}

func AddFollowing(connection *db.DatabaseConnection, user *ent.User, newFollowingUser *ent.User) (*ent.User, error) {
	return user.Update().AddFollowing(newFollowingUser).Save(*connection.Ctx)
}

func RemoveFollowing(connection *db.DatabaseConnection, user *ent.User, target *ent.User) (*ent.User, error) {
	return user.Update().RemoveFollowing(target).Save(*connection.Ctx)
}
