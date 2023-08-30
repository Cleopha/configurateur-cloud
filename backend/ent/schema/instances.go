package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Instances holds the schema definition for the Instances entity.
type Instances struct {
	ent.Schema
}

// Todo: Replace type by an enum

// Fields of the Instances.
func (Instances) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.String("Name"),
		field.Int("CPU"),
		field.Int("Ram_Go"),
		field.Int("Stockage_Go"),
		field.String("GPU"),
		field.String("Type"),
		field.Float("Price"),
	}
}

// Edges of the Instances.
func (Instances) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panierInstances", PanierInstances.Type).
			StorageKey(edge.Column("instance_id")),
	}
}
