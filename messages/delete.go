package messages

import (
	"moments/db"
	"moments/ent/message"
)

func Delete(conn *db.DatabaseConnection, messageID int) (int, error) {
	return conn.Client.Message.Delete().Where(message.IDEQ(messageID)).Exec(*conn.Ctx)
}
