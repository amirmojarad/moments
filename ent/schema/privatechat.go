package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PrivateChat holds the schema definition for the PrivateChat entity.
type PrivateChat struct {
	ent.Schema
}

func (PrivateChat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

// Fields of the PrivateChat.
func (PrivateChat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("receiver_id"),
		field.Int("sender_id"),
	}
}

// Edges of the PrivateChat.
func (PrivateChat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).Ref("sender_pv_chat").Required().Unique().Field("sender_id"),
		edge.From("receiver", User.Type).Ref("receiver_pv_chat").Required().Unique().Field("receiver_id"),
		edge.To("chats", Message.Type),
	}
}

func (PrivateChat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("receiver_id", "sender_id").Unique(),
		index.Fields("sender_id", "receiver_id").Unique(),
	}
}
