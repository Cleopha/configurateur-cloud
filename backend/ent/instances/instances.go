// Code generated by ent, DO NOT EDIT.

package instances

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the instances type in the database.
	Label = "instances"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCPU holds the string denoting the cpu field in the database.
	FieldCPU = "cpu"
	// FieldRAMGo holds the string denoting the ram_go field in the database.
	FieldRAMGo = "ram_go"
	// FieldStockageGo holds the string denoting the stockage_go field in the database.
	FieldStockageGo = "stockage_go"
	// FieldGPU holds the string denoting the gpu field in the database.
	FieldGPU = "gpu"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgePanierInstances holds the string denoting the panierinstances edge name in mutations.
	EdgePanierInstances = "panierInstances"
	// Table holds the table name of the instances in the database.
	Table = "instances"
	// PanierInstancesTable is the table that holds the panierInstances relation/edge.
	PanierInstancesTable = "panier_instances"
	// PanierInstancesInverseTable is the table name for the PanierInstances entity.
	// It exists in this package in order to avoid circular dependency with the "panierinstances" package.
	PanierInstancesInverseTable = "panier_instances"
	// PanierInstancesColumn is the table column denoting the panierInstances relation/edge.
	PanierInstancesColumn = "instance_id"
)

// Columns holds all SQL columns for instances fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCPU,
	FieldRAMGo,
	FieldStockageGo,
	FieldGPU,
	FieldType,
	FieldPrice,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Instances queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCPU orders the results by the CPU field.
func ByCPU(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCPU, opts...).ToFunc()
}

// ByRAMGo orders the results by the Ram_Go field.
func ByRAMGo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRAMGo, opts...).ToFunc()
}

// ByStockageGo orders the results by the Stockage_Go field.
func ByStockageGo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockageGo, opts...).ToFunc()
}

// ByGPU orders the results by the GPU field.
func ByGPU(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGPU, opts...).ToFunc()
}

// ByType orders the results by the Type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPrice orders the results by the Price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByPanierInstancesCount orders the results by panierInstances count.
func ByPanierInstancesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPanierInstancesStep(), opts...)
	}
}

// ByPanierInstances orders the results by panierInstances terms.
func ByPanierInstances(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPanierInstancesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPanierInstancesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PanierInstancesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PanierInstancesTable, PanierInstancesColumn),
	)
}
