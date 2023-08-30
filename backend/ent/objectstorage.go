// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/objectstorage"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ObjectStorage is the model entity for the ObjectStorage schema.
type ObjectStorage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Price holds the value of the "Price" field.
	Price float64 `json:"Price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ObjectStorageQuery when eager-loading is set.
	Edges        ObjectStorageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ObjectStorageEdges holds the relations/edges for other nodes in the graph.
type ObjectStorageEdges struct {
	// PanierObjectStorage holds the value of the panierObjectStorage edge.
	PanierObjectStorage []*PanierObjectStorage `json:"panierObjectStorage,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PanierObjectStorageOrErr returns the PanierObjectStorage value or an error if the edge
// was not loaded in eager-loading.
func (e ObjectStorageEdges) PanierObjectStorageOrErr() ([]*PanierObjectStorage, error) {
	if e.loadedTypes[0] {
		return e.PanierObjectStorage, nil
	}
	return nil, &NotLoadedError{edge: "panierObjectStorage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ObjectStorage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case objectstorage.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case objectstorage.FieldID:
			values[i] = new(sql.NullInt64)
		case objectstorage.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ObjectStorage fields.
func (os *ObjectStorage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case objectstorage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			os.ID = int(value.Int64)
		case objectstorage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				os.Name = value.String
			}
		case objectstorage.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Price", values[i])
			} else if value.Valid {
				os.Price = value.Float64
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ObjectStorage.
// This includes values selected through modifiers, order, etc.
func (os *ObjectStorage) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryPanierObjectStorage queries the "panierObjectStorage" edge of the ObjectStorage entity.
func (os *ObjectStorage) QueryPanierObjectStorage() *PanierObjectStorageQuery {
	return NewObjectStorageClient(os.config).QueryPanierObjectStorage(os)
}

// Update returns a builder for updating this ObjectStorage.
// Note that you need to call ObjectStorage.Unwrap() before calling this method if this ObjectStorage
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *ObjectStorage) Update() *ObjectStorageUpdateOne {
	return NewObjectStorageClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the ObjectStorage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *ObjectStorage) Unwrap() *ObjectStorage {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: ObjectStorage is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *ObjectStorage) String() string {
	var builder strings.Builder
	builder.WriteString("ObjectStorage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("Name=")
	builder.WriteString(os.Name)
	builder.WriteString(", ")
	builder.WriteString("Price=")
	builder.WriteString(fmt.Sprintf("%v", os.Price))
	builder.WriteByte(')')
	return builder.String()
}

// ObjectStorages is a parsable slice of ObjectStorage.
type ObjectStorages []*ObjectStorage