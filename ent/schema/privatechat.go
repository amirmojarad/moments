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

// Fields of the PrivateChat.
func (PrivateChat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("first_user_id").Optional(),
		field.Int("second_user_id").Optional(),
	}

}

// Edges of the PrivateChat.
func (PrivateChat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("first_user", User.Type).Ref("my_pv_chats").Unique().Field("first_user_id"),
		edge.From("second_user", User.Type).Ref("other_pv_chats").Unique().Field("second_user_id"),
	}
}

func (PrivateChat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("first_user_id", "second_user_id").Unique(),
	}
}
