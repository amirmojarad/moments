package schema

import "entgo.io/ent"

// PrivateChat holds the schema definition for the PrivateChat entity.
type PrivateChat struct {
	ent.Schema
}

// Fields of the PrivateChat.
func (PrivateChat) Fields() []ent.Field {
	return nil
}

// Edges of the PrivateChat.
func (PrivateChat) Edges() []ent.Edge {
	return nil
}
