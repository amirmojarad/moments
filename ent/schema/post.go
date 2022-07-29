package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Text("plain_text"),
		field.Uint64("likes").Default(0).Optional(),
		field.String("image_url").Optional(),
		field.String("attached_file_url").Optional(),
		field.String("link_url").Optional(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("posts").Unique(),
	}
}
