package messages

import (
	"moments/db"
	"moments/ent"
	"time"
)

func ChangeMessageText(conn *db.DatabaseConnection, message *ent.Message, text string) (*ent.Message, error) {
	return message.Update().SetText(text).SetUpdatedDate(time.Now()).Save(*conn.Ctx)
}
