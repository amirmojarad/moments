package room

import (
	"moments/db"
	"moments/ent/room"
)

func DeleteRoomByID(conn *db.DatabaseConnection, roomID int) (int, error) {
	return conn.Client.Room.Delete().Where(room.IDEQ(roomID)).Exec(*conn.Ctx)
}
