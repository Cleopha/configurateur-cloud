// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/ipaddress"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// IPAddress is the model entity for the IPAddress schema.
type IPAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Price holds the value of the "Price" field.
	Price float64 `json:"Price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IPAddressQuery when eager-loading is set.
	Edges        IPAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// IPAddressEdges holds the relations/edges for other nodes in the graph.
type IPAddressEdges struct {
	// PanierIPAddress holds the value of the panierIPAddress edge.
	PanierIPAddress []*PanierIPAddress `json:"panierIPAddress,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PanierIPAddressOrErr returns the PanierIPAddress value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) PanierIPAddressOrErr() ([]*PanierIPAddress, error) {
	if e.loadedTypes[0] {
		return e.PanierIPAddress, nil
	}
	return nil, &NotLoadedError{edge: "panierIPAddress"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IPAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case ipaddress.FieldID:
			values[i] = new(sql.NullInt64)
		case ipaddress.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IPAddress fields.
func (ia *IPAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ia.ID = int(value.Int64)
		case ipaddress.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				ia.Name = value.String
			}
		case ipaddress.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Price", values[i])
			} else if value.Valid {
				ia.Price = value.Float64
			}
		default:
			ia.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IPAddress.
// This includes values selected through modifiers, order, etc.
func (ia *IPAddress) Value(name string) (ent.Value, error) {
	return ia.selectValues.Get(name)
}

// QueryPanierIPAddress queries the "panierIPAddress" edge of the IPAddress entity.
func (ia *IPAddress) QueryPanierIPAddress() *PanierIPAddressQuery {
	return NewIPAddressClient(ia.config).QueryPanierIPAddress(ia)
}

// Update returns a builder for updating this IPAddress.
// Note that you need to call IPAddress.Unwrap() before calling this method if this IPAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ia *IPAddress) Update() *IPAddressUpdateOne {
	return NewIPAddressClient(ia.config).UpdateOne(ia)
}

// Unwrap unwraps the IPAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ia *IPAddress) Unwrap() *IPAddress {
	_tx, ok := ia.config.driver.(*txDriver)
	if !ok {
		panic("ent: IPAddress is not a transactional entity")
	}
	ia.config.driver = _tx.drv
	return ia
}

// String implements the fmt.Stringer.
func (ia *IPAddress) String() string {
	var builder strings.Builder
	builder.WriteString("IPAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ia.ID))
	builder.WriteString("Name=")
	builder.WriteString(ia.Name)
	builder.WriteString(", ")
	builder.WriteString("Price=")
	builder.WriteString(fmt.Sprintf("%v", ia.Price))
	builder.WriteByte(')')
	return builder.String()
}

// IPAddresses is a parsable slice of IPAddress.
type IPAddresses []*IPAddress
