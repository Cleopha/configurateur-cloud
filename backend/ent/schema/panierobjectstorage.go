package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PanierObjectStorage holds the schema definition for the PanierObjectStorage entity.
type PanierObjectStorage struct {
	ent.Schema
}

// Fields of the PanierObjectStorage.
func (PanierObjectStorage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the PanierObjectStorage.
func (PanierObjectStorage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("panier", Panier.Type).
			Ref("panierObjectStorage").
			Unique().
			Required(),
		edge.From("objectStorage", ObjectStorage.Type).
			Ref("panierObjectStorage").
			Unique().
			Required(),
	}
}
