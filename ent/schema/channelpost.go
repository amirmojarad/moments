package schema

import "entgo.io/ent"

// ChannelPost holds the schema definition for the ChannelPost entity.
type ChannelPost struct {
	ent.Schema
}

// Fields of the ChannelPost.
func (ChannelPost) Fields() []ent.Field {
	return nil
}

// Edges of the ChannelPost.
func (ChannelPost) Edges() []ent.Edge {
	return nil
}
