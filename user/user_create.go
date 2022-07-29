package user

import (
	"github.com/rs/zerolog/log"
	"moments/db"
	"moments/ent"
	"moments/utils"
)

// CreateUser create and return a new user
func CreateUser(db *db.DatabaseConnection, user *ent.User) (*ent.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Error().Err(err).Msg("error while hashing password")
		return nil, err
	}
	// isAdmin and isStaff is false by default
	return db.Client.User.Create().
		SetEmail(user.Email).
		SetPassword(hashedPassword).
		SetUsername(user.Username).
		Save(*db.Ctx)
}

func CreateSuperUser(db *db.DatabaseConnection, user *ent.User) (*ent.User, error) {
	createdUser, err := CreateUser(db, user)
	if err != nil {
		log.Error().Err(err).Msg("error while creating a new user in database")
		return nil, err
	}
	return createdUser.Update().SetIsAdmin(true).SetIsStaff(true).Save(*db.Ctx)
}
