package user

import (
	"errors"
	"moments/db"
	"moments/ent"
	"moments/ent/user"
	"moments/utils"
	"time"
)

func ChangePassword(conn *db.DatabaseConnection, username string, newPassword, currentPassword string) (*ent.User, error) {
	userByUsername, err := GetUserByUsername(conn, username)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(currentPassword, userByUsername.Password) {
		return nil, errors.New("wrong current password")
	}
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}
	return userByUsername.Update().SetPassword(hashedPassword).SetUpdatedDate(time.Now()).Save(*conn.Ctx)
}

func UpdateUserWithUsername(conn *db.DatabaseConnection, user *ent.User, username string) (*ent.User, error) {
	userByUsername, err := GetUserByUsername(conn, username)
	if err != nil {
		return nil, err
	}
	return userByUsername.Update().SetNillableFullName(user.FullName).SetUpdatedDate(time.Now()).Save(*conn.Ctx)
}

func AddFollowing(connection *db.DatabaseConnection, user *ent.User, newFollowingUser *ent.User) (*ent.User, error) {
	return user.Update().AddFollowing(newFollowingUser).Save(*connection.Ctx)
}

func RemoveFollowing(connection *db.DatabaseConnection, user *ent.User, target *ent.User) (*ent.User, error) {
	return user.Update().RemoveFollowing(target).Save(*connection.Ctx)
}

func GetAllUsersByUsernames(conn *db.DatabaseConnection, usernames ...string) ([]int, error) {
	return conn.Client.User.Query().Where(user.UsernameIn(usernames...)).IDs(*conn.Ctx)
}

func RemoveFollowings(conn *db.DatabaseConnection, u *ent.User, targetsUsernames ...string) (*ent.User, error) {
	usersIDs, err := GetAllUsersByUsernames(conn, targetsUsernames...)
	if err != nil {
		return nil, err
	}
	return u.Update().RemoveFollowingIDs(usersIDs...).Save(*conn.Ctx)
}
