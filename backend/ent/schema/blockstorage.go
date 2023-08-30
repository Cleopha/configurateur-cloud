package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BlockStorage holds the schema definition for the BlockStorage entity.
type BlockStorage struct {
	ent.Schema
}

// Fields of the BlockStorage.
func (BlockStorage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.String("Name"),
		field.Int("IOPS"),
		field.Float("Bandwidth"),
		field.Float("Price"),
	}
}

// Edges of the BlockStorage.
func (BlockStorage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panierBlockStorage", PanierBlockStorage.Type),
	}
}
