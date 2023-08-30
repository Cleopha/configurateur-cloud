package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PanierIPAddress holds the schema definition for the PanierIPAddress entity.
type PanierIPAddress struct {
	ent.Schema
}

// Fields of the PanierIPAddress.
func (PanierIPAddress) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the PanierIPAddress.
func (PanierIPAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("panier", Panier.Type).
			Ref("panierIPAddress").
			Unique().
			Required(),
		edge.From("IPAddress", IPAddress.Type).
			Ref("panierIPAddress").
			Unique().
			Required(),
	}
}
