package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"net/mail"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty().MinLen(4).MaxLen(255),
		field.String("email").Unique().NotEmpty().Validate(validateEmail).MaxLen(255),
		field.String("password").NotEmpty().MaxLen(255),
		field.String("full_name").Nillable().Optional().MaxLen(255),
		field.Bool("is_admin").Default(false).Optional(),
		field.Bool("is_staff").Default(false).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("following", User.Type).From("followers"),
	}
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
