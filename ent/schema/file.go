package schema

import "entgo.io/ent"

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// TODO create a model to store information about stored file in minio

// Fields of the File.
func (File) Fields() []ent.Field {
	return nil
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
