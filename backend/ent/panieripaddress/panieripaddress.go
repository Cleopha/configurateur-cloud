// Code generated by ent, DO NOT EDIT.

package panieripaddress

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the panieripaddress type in the database.
	Label = "panier_ip_address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// EdgePanier holds the string denoting the panier edge name in mutations.
	EdgePanier = "panier"
	// EdgeIPAddress holds the string denoting the ipaddress edge name in mutations.
	EdgeIPAddress = "IPAddress"
	// Table holds the table name of the panieripaddress in the database.
	Table = "panier_ip_addresses"
	// PanierTable is the table that holds the panier relation/edge.
	PanierTable = "panier_ip_addresses"
	// PanierInverseTable is the table name for the Panier entity.
	// It exists in this package in order to avoid circular dependency with the "panier" package.
	PanierInverseTable = "paniers"
	// PanierColumn is the table column denoting the panier relation/edge.
	PanierColumn = "panier_panier_ip_address"
	// IPAddressTable is the table that holds the IPAddress relation/edge.
	IPAddressTable = "panier_ip_addresses"
	// IPAddressInverseTable is the table name for the IPAddress entity.
	// It exists in this package in order to avoid circular dependency with the "ipaddress" package.
	IPAddressInverseTable = "ip_addresses"
	// IPAddressColumn is the table column denoting the IPAddress relation/edge.
	IPAddressColumn = "IPAddress_id"
)

// Columns holds all SQL columns for panieripaddress fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "panier_ip_addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"IPAddress_id",
	"panier_panier_ip_address",
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

// OrderOption defines the ordering options for the PanierIPAddress queries.
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

// ByIPAddressField orders the results by IPAddress field.
func ByIPAddressField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIPAddressStep(), sql.OrderByField(field, opts...))
	}
}
func newPanierStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PanierInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PanierTable, PanierColumn),
	)
}
func newIPAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IPAddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, IPAddressTable, IPAddressColumn),
	)
}
