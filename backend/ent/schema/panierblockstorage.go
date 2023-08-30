package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PanierBlockStorage holds the schema definition for the PanierBlockStorage entity.
type PanierBlockStorage struct {
	ent.Schema
}

// Fields of the PanierBlockStorage.
func (PanierBlockStorage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the PanierBlockStorage.
func (PanierBlockStorage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("panier", Panier.Type).
			Ref("panierBlockStorage").
			Unique().
			Required(),
		edge.From("blockStorage", BlockStorage.Type).
			Ref("panierBlockStorage").
			Unique().
			Required(),
	}
}
