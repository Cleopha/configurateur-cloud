// Code generated by ent, DO NOT EDIT.

package panierinstances

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the panierinstances type in the database.
	Label = "panier_instances"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// EdgePanier holds the string denoting the panier edge name in mutations.
	EdgePanier = "panier"
	// EdgeInstance holds the string denoting the instance edge name in mutations.
	EdgeInstance = "instance"
	// Table holds the table name of the panierinstances in the database.
	Table = "panier_instances"
	// PanierTable is the table that holds the panier relation/edge.
	PanierTable = "panier_instances"
	// PanierInverseTable is the table name for the Panier entity.
	// It exists in this package in order to avoid circular dependency with the "panier" package.
	PanierInverseTable = "paniers"
	// PanierColumn is the table column denoting the panier relation/edge.
	PanierColumn = "panier_panier_instances"
	// InstanceTable is the table that holds the instance relation/edge.
	InstanceTable = "panier_instances"
	// InstanceInverseTable is the table name for the Instances entity.
	// It exists in this package in order to avoid circular dependency with the "instances" package.
	InstanceInverseTable = "instances"
	// InstanceColumn is the table column denoting the instance relation/edge.
	InstanceColumn = "instance_id"
)

// Columns holds all SQL columns for panierinstances fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "panier_instances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"instance_id",
	"panier_panier_instances",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the PanierInstances queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByPanierField orders the results by panier field.
func ByPanierField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPanierStep(), sql.OrderByField(field, opts...))
	}
}

// ByInstanceField orders the results by instance field.
func ByInstanceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstanceStep(), sql.OrderByField(field, opts...))
	}
}
func newPanierStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PanierInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PanierTable, PanierColumn),
	)
}
func newInstanceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstanceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
	)
}