package room

import (
	"github.com/rs/zerolog/log"
	"moments/db"
	"moments/ent"
	"moments/ent/room"
	"moments/ent/user"
)

func GetRoomByID(connection *db.DatabaseConnection, roomID int) (*ent.Room, error) {
	return connection.Client.Room.Get(*connection.Ctx, roomID)
}

func GetUserPublicRoomByTitle(connection *db.DatabaseConnection, user *ent.User, title string) (*ent.Room, error) {
	return user.QueryRooms().Where(room.TypeEQ(room.TypeGroup)).Where(room.TitleEQ(title)).Only(*connection.Ctx)
}

func GetUserPublicRoomByID(connection *db.DatabaseConnection, user *ent.User, id int) (*ent.Room, error) {
	return user.QueryRooms().Where(room.TypeEQ(room.TypeGroup)).Where(room.IDEQ(id)).Only(*connection.Ctx)
}

func GetAllUserPublicRooms(connection *db.DatabaseConnection, user *ent.User) ([]*ent.Room, error) {
	return user.QueryRooms().Where(room.TypeEQ(room.TypeGroup)).All(*connection.Ctx)
}

func GetAllUserPrivateRooms(connection *db.DatabaseConnection, user *ent.User) ([]*ent.Room, error) {
	return user.QueryRooms().Where(room.TypeEQ(room.TypePrivate)).All(*connection.Ctx)
}

func GetUserPrivateRoomByID(connection *db.DatabaseConnection, user *ent.User, id int) (*ent.Room, error) {
	return user.QueryRooms().Where(room.TypeEQ(room.TypePrivate)).Where(room.IDEQ(id)).Only(*connection.Ctx)
}

func GetPrivateUserParticipants(connection *db.DatabaseConnection, id int) ([]*ent.User, error) {
	roomByID, err := GetRoomByID(connection, id)
	if err != nil {
		log.Error().Err(err).Msg("error while getting room from database in GetPrivateUserParticipants function")
		return nil, err
	}
	return roomByID.QueryUsers().All(*connection.Ctx)
}

func IsUserInRoom(conn *db.DatabaseConnection, u *ent.User, roomID int) (bool, error) {
	roomByID, err := GetRoomByID(conn, roomID)
	if err != nil {
		return false, err
	}
	_, err = roomByID.QueryUsers().Where(user.UsernameEQ(u.Username)).First(*conn.Ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
