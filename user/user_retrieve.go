package user

import (
	"moments/db"
	"moments/ent"
	"moments/ent/user"
)

func GetUserByID(connection *db.DatabaseConnection, id int) (*ent.User, error) {
	return connection.Client.User.Get(*connection.Ctx, id)
}

func GetUserByUsername(connection *db.DatabaseConnection, username string) (*ent.User, error) {
	return connection.Client.User.Query().Where(user.UsernameEQ(username)).First(*connection.Ctx)
}

func GetAllFollowers(connection *db.DatabaseConnection, user *ent.User) ([]*ent.User, error) {
	return user.QueryFollowers().All(*connection.Ctx)
}

func GetAllFollowing(connection *db.DatabaseConnection, user *ent.User) ([]*ent.User, error) {
	return user.QueryFollowing().All(*connection.Ctx)
}

func IsFollowing(connection *db.DatabaseConnection, u *ent.User, targetUser *ent.User) (*ent.User, error) {
	return u.QueryFollowing().Where(user.UsernameEQ(targetUser.Username)).Only(*connection.Ctx)
}

func IsFollower(connection *db.DatabaseConnection, u *ent.User, targetUser *ent.User) (*ent.User, error) {
	return u.QueryFollowers().Where(user.UsernameEQ(targetUser.Username)).Only(*connection.Ctx)
}
