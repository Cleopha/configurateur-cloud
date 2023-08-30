package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Panier holds the schema definition for the Panier entity.
type Panier struct {
	ent.Schema
}

// Fields of the Panier.
func (Panier) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uid", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
	}
}

// Edges of the Panier.
func (Panier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panierInstances", PanierInstances.Type),
		edge.To("panierBlockStorage", PanierBlockStorage.Type),
		edge.To("panierIPAddress", PanierIPAddress.Type),
		edge.To("panierLoadBalancer", PanierLoadBalancer.Type),
		edge.To("panierObjectStorage", PanierObjectStorage.Type),
	}
}
