// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/ipaddress"
	"backend/ent/panier"
	"backend/ent/panieripaddress"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierIPAddressUpdate is the builder for updating PanierIPAddress entities.
type PanierIPAddressUpdate struct {
	config
	hooks    []Hook
	mutation *PanierIPAddressMutation
}

// Where appends a list predicates to the PanierIPAddressUpdate builder.
func (piau *PanierIPAddressUpdate) Where(ps ...predicate.PanierIPAddress) *PanierIPAddressUpdate {
	piau.mutation.Where(ps...)
	return piau
}

// SetQuantity sets the "quantity" field.
func (piau *PanierIPAddressUpdate) SetQuantity(i int) *PanierIPAddressUpdate {
	piau.mutation.ResetQuantity()
	piau.mutation.SetQuantity(i)
	return piau
}

// AddQuantity adds i to the "quantity" field.
func (piau *PanierIPAddressUpdate) AddQuantity(i int) *PanierIPAddressUpdate {
	piau.mutation.AddQuantity(i)
	return piau
}

// SetPanierID sets the "panier" edge to the Panier entity by ID.
func (piau *PanierIPAddressUpdate) SetPanierID(id int) *PanierIPAddressUpdate {
	piau.mutation.SetPanierID(id)
	return piau
}

// SetPanier sets the "panier" edge to the Panier entity.
func (piau *PanierIPAddressUpdate) SetPanier(p *Panier) *PanierIPAddressUpdate {
	return piau.SetPanierID(p.ID)
}

// SetIPAddressID sets the "IPAddress" edge to the IPAddress entity by ID.
func (piau *PanierIPAddressUpdate) SetIPAddressID(id int) *PanierIPAddressUpdate {
	piau.mutation.SetIPAddressID(id)
	return piau
}

// SetIPAddress sets the "IPAddress" edge to the IPAddress entity.
func (piau *PanierIPAddressUpdate) SetIPAddress(i *IPAddress) *PanierIPAddressUpdate {
	return piau.SetIPAddressID(i.ID)
}

// Mutation returns the PanierIPAddressMutation object of the builder.
func (piau *PanierIPAddressUpdate) Mutation() *PanierIPAddressMutation {
	return piau.mutation
}

// ClearPanier clears the "panier" edge to the Panier entity.
func (piau *PanierIPAddressUpdate) ClearPanier() *PanierIPAddressUpdate {
	piau.mutation.ClearPanier()
	return piau
}

// ClearIPAddress clears the "IPAddress" edge to the IPAddress entity.
func (piau *PanierIPAddressUpdate) ClearIPAddress() *PanierIPAddressUpdate {
	piau.mutation.ClearIPAddress()
	return piau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piau *PanierIPAddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, piau.sqlSave, piau.mutation, piau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piau *PanierIPAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := piau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piau *PanierIPAddressUpdate) Exec(ctx context.Context) error {
	_, err := piau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piau *PanierIPAddressUpdate) ExecX(ctx context.Context) {
	if err := piau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piau *PanierIPAddressUpdate) check() error {
	if _, ok := piau.mutation.PanierID(); piau.mutation.PanierCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierIPAddress.panier"`)
	}
	if _, ok := piau.mutation.IPAddressID(); piau.mutation.IPAddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierIPAddress.IPAddress"`)
	}
	return nil
}

func (piau *PanierIPAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(panieripaddress.Table, panieripaddress.Columns, sqlgraph.NewFieldSpec(panieripaddress.FieldID, field.TypeInt))
	if ps := piau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piau.mutation.Quantity(); ok {
		_spec.SetField(panieripaddress.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := piau.mutation.AddedQuantity(); ok {
		_spec.AddField(panieripaddress.FieldQuantity, field.TypeInt, value)
	}
	if piau.mutation.PanierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.PanierTable,
			Columns: []string{panieripaddress.PanierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panier.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piau.mutation.PanierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.PanierTable,
			Columns: []string{panieripaddress.PanierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panier.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piau.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.IPAddressTable,
			Columns: []string{panieripaddress.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piau.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.IPAddressTable,
			Columns: []string{panieripaddress.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{panieripaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piau.mutation.done = true
	return n, nil
}

// PanierIPAddressUpdateOne is the builder for updating a single PanierIPAddress entity.
type PanierIPAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PanierIPAddressMutation
}

// SetQuantity sets the "quantity" field.
func (piauo *PanierIPAddressUpdateOne) SetQuantity(i int) *PanierIPAddressUpdateOne {
	piauo.mutation.ResetQuantity()
	piauo.mutation.SetQuantity(i)
	return piauo
}

// AddQuantity adds i to the "quantity" field.
func (piauo *PanierIPAddressUpdateOne) AddQuantity(i int) *PanierIPAddressUpdateOne {
	piauo.mutation.AddQuantity(i)
	return piauo
}

// SetPanierID sets the "panier" edge to the Panier entity by ID.
func (piauo *PanierIPAddressUpdateOne) SetPanierID(id int) *PanierIPAddressUpdateOne {
	piauo.mutation.SetPanierID(id)
	return piauo
}

// SetPanier sets the "panier" edge to the Panier entity.
func (piauo *PanierIPAddressUpdateOne) SetPanier(p *Panier) *PanierIPAddressUpdateOne {
	return piauo.SetPanierID(p.ID)
}

// SetIPAddressID sets the "IPAddress" edge to the IPAddress entity by ID.
func (piauo *PanierIPAddressUpdateOne) SetIPAddressID(id int) *PanierIPAddressUpdateOne {
	piauo.mutation.SetIPAddressID(id)
	return piauo
}

// SetIPAddress sets the "IPAddress" edge to the IPAddress entity.
func (piauo *PanierIPAddressUpdateOne) SetIPAddress(i *IPAddress) *PanierIPAddressUpdateOne {
	return piauo.SetIPAddressID(i.ID)
}

// Mutation returns the PanierIPAddressMutation object of the builder.
func (piauo *PanierIPAddressUpdateOne) Mutation() *PanierIPAddressMutation {
	return piauo.mutation
}

// ClearPanier clears the "panier" edge to the Panier entity.
func (piauo *PanierIPAddressUpdateOne) ClearPanier() *PanierIPAddressUpdateOne {
	piauo.mutation.ClearPanier()
	return piauo
}

// ClearIPAddress clears the "IPAddress" edge to the IPAddress entity.
func (piauo *PanierIPAddressUpdateOne) ClearIPAddress() *PanierIPAddressUpdateOne {
	piauo.mutation.ClearIPAddress()
	return piauo
}

// Where appends a list predicates to the PanierIPAddressUpdate builder.
func (piauo *PanierIPAddressUpdateOne) Where(ps ...predicate.PanierIPAddress) *PanierIPAddressUpdateOne {
	piauo.mutation.Where(ps...)
	return piauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piauo *PanierIPAddressUpdateOne) Select(field string, fields ...string) *PanierIPAddressUpdateOne {
	piauo.fields = append([]string{field}, fields...)
	return piauo
}

// Save executes the query and returns the updated PanierIPAddress entity.
func (piauo *PanierIPAddressUpdateOne) Save(ctx context.Context) (*PanierIPAddress, error) {
	return withHooks(ctx, piauo.sqlSave, piauo.mutation, piauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piauo *PanierIPAddressUpdateOne) SaveX(ctx context.Context) *PanierIPAddress {
	node, err := piauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piauo *PanierIPAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := piauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piauo *PanierIPAddressUpdateOne) ExecX(ctx context.Context) {
	if err := piauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piauo *PanierIPAddressUpdateOne) check() error {
	if _, ok := piauo.mutation.PanierID(); piauo.mutation.PanierCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierIPAddress.panier"`)
	}
	if _, ok := piauo.mutation.IPAddressID(); piauo.mutation.IPAddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierIPAddress.IPAddress"`)
	}
	return nil
}

func (piauo *PanierIPAddressUpdateOne) sqlSave(ctx context.Context) (_node *PanierIPAddress, err error) {
	if err := piauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(panieripaddress.Table, panieripaddress.Columns, sqlgraph.NewFieldSpec(panieripaddress.FieldID, field.TypeInt))
	id, ok := piauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PanierIPAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, panieripaddress.FieldID)
		for _, f := range fields {
			if !panieripaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != panieripaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piauo.mutation.Quantity(); ok {
		_spec.SetField(panieripaddress.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := piauo.mutation.AddedQuantity(); ok {
		_spec.AddField(panieripaddress.FieldQuantity, field.TypeInt, value)
	}
	if piauo.mutation.PanierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.PanierTable,
			Columns: []string{panieripaddress.PanierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panier.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piauo.mutation.PanierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.PanierTable,
			Columns: []string{panieripaddress.PanierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panier.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piauo.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.IPAddressTable,
			Columns: []string{panieripaddress.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piauo.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panieripaddress.IPAddressTable,
			Columns: []string{panieripaddress.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PanierIPAddress{config: piauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{panieripaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piauo.mutation.done = true
	return _node, nil
}
