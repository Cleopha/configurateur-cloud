package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LoadBalancer holds the schema definition for the LoadBalancer entity.
type LoadBalancer struct {
	ent.Schema
}

// Fields of the LoadBalancer.
func (LoadBalancer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.String("Name"),
		field.Float("Price"),
	}
}

// Edges of the LoadBalancer.
func (LoadBalancer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panierLoadBalancer", PanierLoadBalancer.Type).
			StorageKey(edge.Column("loadBalancer_id")),
	}
}
