package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PanierInstances holds the schema definition for the PanierInstances entity.
type PanierInstances struct {
	ent.Schema
}

// Fields of the PanierInstances.
func (PanierInstances) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the PanierInstances.
func (PanierInstances) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("panier", Panier.Type).
			Ref("panierInstances").
			Unique().
			Required(),
		edge.From("instance", Instances.Type).
			Ref("panierInstances").
			Unique().
			Required(),
	}
}
