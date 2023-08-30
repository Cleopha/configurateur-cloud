package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ObjectStorage holds the schema definition for the ObjectStorage entity.
type ObjectStorage struct {
	ent.Schema
}

// Fields of the ObjectStorage.
func (ObjectStorage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.String("Name"),
		field.Float("Price"),
	}
}

// Edges of the ObjectStorage.
func (ObjectStorage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panierObjectStorage", PanierObjectStorage.Type).
			StorageKey(edge.Column("objectStorage_id")),
	}
}
