package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PanierLoadBalancer holds the schema definition for the PanierLoadBalancer entity.
type PanierLoadBalancer struct {
	ent.Schema
}

// Fields of the PanierLoadBalancer.
func (PanierLoadBalancer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the PanierLoadBalancer.
func (PanierLoadBalancer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("panier", Panier.Type).
			Ref("panierLoadBalancer").
			Unique().
			Required(),
		edge.From("loadBalancer", LoadBalancer.Type).
			Ref("panierLoadBalancer").
			Unique().
			Required(),
	}
}
