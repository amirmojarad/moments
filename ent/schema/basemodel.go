package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// BaseModel holds the schema definition for the BaseModel entity.
type BaseModel struct {
	mixin.Schema
}

// Fields of the BaseModel.
func (BaseModel) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_date").Default(time.Now()),
		field.Time("updated_date").Default(time.Now()),
		field.Time("deleted_date").Nillable().Optional(),
	}
}

// Edges of the BaseModel.
func (BaseModel) Edges() []ent.Edge {
	return nil
}
