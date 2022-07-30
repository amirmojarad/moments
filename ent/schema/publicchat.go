package schema

import "entgo.io/ent"

// PublicChat holds the schema definition for the PublicChat entity.
type PublicChat struct {
	ent.Schema
}

// Fields of the PublicChat.
func (PublicChat) Fields() []ent.Field {
	return nil
}

// Edges of the PublicChat.
func (PublicChat) Edges() []ent.Edge {
	return nil
}
