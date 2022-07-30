package privateChat

import (
	"github.com/rs/zerolog/log"
	"moments/db"
	"moments/ent"
	"moments/ent/privatechat"
)

func GetPrivateChatBySender(connection *db.DatabaseConnection, sender *ent.User) ([]*ent.PrivateChat, error) {
	return sender.QuerySenderPvChat().All(*connection.Ctx)
}

func GetPrivateChatByID(connection *db.DatabaseConnection, u *ent.User, pvChatID int) (*ent.PrivateChat, error) {
	receiver, err := u.QueryReceiverPvChat().Where(privatechat.IDEQ(pvChatID)).Only(*connection.Ctx)
	if err != nil {
		log.Error().Err(err).Msg("error while fetching receiver in private chats")
		return nil, err
	}
	if receiver != nil {
		return receiver, nil
	}

	sender, err := u.QuerySenderPvChat().Where(privatechat.IDEQ(pvChatID)).Only(*connection.Ctx)
	if err != nil {
		log.Error().Err(err).Msg("error while fetching sender in private chats")
		return sender, err
	}
	if receiver != nil {
		return sender, nil
	}
	return nil, err
}
