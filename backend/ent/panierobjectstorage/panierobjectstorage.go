// Code generated by ent, DO NOT EDIT.

package panierobjectstorage

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the panierobjectstorage type in the database.
	Label = "panier_object_storage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// EdgePanier holds the string denoting the panier edge name in mutations.
	EdgePanier = "panier"
	// EdgeObjectStorage holds the string denoting the objectstorage edge name in mutations.
	EdgeObjectStorage = "objectStorage"
	// Table holds the table name of the panierobjectstorage in the database.
	Table = "panier_object_storages"
	// PanierTable is the table that holds the panier relation/edge.
	PanierTable = "panier_object_storages"
	// PanierInverseTable is the table name for the Panier entity.
	// It exists in this package in order to avoid circular dependency with the "panier" package.
	PanierInverseTable = "paniers"
	// PanierColumn is the table column denoting the panier relation/edge.
	PanierColumn = "panier_panier_object_storage"
	// ObjectStorageTable is the table that holds the objectStorage relation/edge.
	ObjectStorageTable = "panier_object_storages"
	// ObjectStorageInverseTable is the table name for the ObjectStorage entity.
	// It exists in this package in order to avoid circular dependency with the "objectstorage" package.
	ObjectStorageInverseTable = "object_storages"
	// ObjectStorageColumn is the table column denoting the objectStorage relation/edge.
	ObjectStorageColumn = "objectStorage_id"
)

// Columns holds all SQL columns for panierobjectstorage fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "panier_object_storages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"objectStorage_id",
	"panier_panier_object_storage",
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

// OrderOption defines the ordering options for the PanierObjectStorage queries.
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

// ByObjectStorageField orders the results by objectStorage field.
func ByObjectStorageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newObjectStorageStep(), sql.OrderByField(field, opts...))
	}
}
func newPanierStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PanierInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PanierTable, PanierColumn),
	)
}
func newObjectStorageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ObjectStorageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ObjectStorageTable, ObjectStorageColumn),
	)
}
