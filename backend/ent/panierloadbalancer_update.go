// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/loadbalancer"
	"backend/ent/panier"
	"backend/ent/panierloadbalancer"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierLoadBalancerUpdate is the builder for updating PanierLoadBalancer entities.
type PanierLoadBalancerUpdate struct {
	config
	hooks    []Hook
	mutation *PanierLoadBalancerMutation
}

// Where appends a list predicates to the PanierLoadBalancerUpdate builder.
func (plbu *PanierLoadBalancerUpdate) Where(ps ...predicate.PanierLoadBalancer) *PanierLoadBalancerUpdate {
	plbu.mutation.Where(ps...)
	return plbu
}

// SetQuantity sets the "quantity" field.
func (plbu *PanierLoadBalancerUpdate) SetQuantity(i int) *PanierLoadBalancerUpdate {
	plbu.mutation.ResetQuantity()
	plbu.mutation.SetQuantity(i)
	return plbu
}

// AddQuantity adds i to the "quantity" field.
func (plbu *PanierLoadBalancerUpdate) AddQuantity(i int) *PanierLoadBalancerUpdate {
	plbu.mutation.AddQuantity(i)
	return plbu
}

// SetPanierID sets the "panier" edge to the Panier entity by ID.
func (plbu *PanierLoadBalancerUpdate) SetPanierID(id int) *PanierLoadBalancerUpdate {
	plbu.mutation.SetPanierID(id)
	return plbu
}

// SetPanier sets the "panier" edge to the Panier entity.
func (plbu *PanierLoadBalancerUpdate) SetPanier(p *Panier) *PanierLoadBalancerUpdate {
	return plbu.SetPanierID(p.ID)
}

// SetLoadBalancerID sets the "loadBalancer" edge to the LoadBalancer entity by ID.
func (plbu *PanierLoadBalancerUpdate) SetLoadBalancerID(id int) *PanierLoadBalancerUpdate {
	plbu.mutation.SetLoadBalancerID(id)
	return plbu
}

// SetLoadBalancer sets the "loadBalancer" edge to the LoadBalancer entity.
func (plbu *PanierLoadBalancerUpdate) SetLoadBalancer(l *LoadBalancer) *PanierLoadBalancerUpdate {
	return plbu.SetLoadBalancerID(l.ID)
}

// Mutation returns the PanierLoadBalancerMutation object of the builder.
func (plbu *PanierLoadBalancerUpdate) Mutation() *PanierLoadBalancerMutation {
	return plbu.mutation
}

// ClearPanier clears the "panier" edge to the Panier entity.
func (plbu *PanierLoadBalancerUpdate) ClearPanier() *PanierLoadBalancerUpdate {
	plbu.mutation.ClearPanier()
	return plbu
}

// ClearLoadBalancer clears the "loadBalancer" edge to the LoadBalancer entity.
func (plbu *PanierLoadBalancerUpdate) ClearLoadBalancer() *PanierLoadBalancerUpdate {
	plbu.mutation.ClearLoadBalancer()
	return plbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (plbu *PanierLoadBalancerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, plbu.sqlSave, plbu.mutation, plbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (plbu *PanierLoadBalancerUpdate) SaveX(ctx context.Context) int {
	affected, err := plbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (plbu *PanierLoadBalancerUpdate) Exec(ctx context.Context) error {
	_, err := plbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plbu *PanierLoadBalancerUpdate) ExecX(ctx context.Context) {
	if err := plbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plbu *PanierLoadBalancerUpdate) check() error {
	if _, ok := plbu.mutation.PanierID(); plbu.mutation.PanierCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierLoadBalancer.panier"`)
	}
	if _, ok := plbu.mutation.LoadBalancerID(); plbu.mutation.LoadBalancerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierLoadBalancer.loadBalancer"`)
	}
	return nil
}

func (plbu *PanierLoadBalancerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := plbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(panierloadbalancer.Table, panierloadbalancer.Columns, sqlgraph.NewFieldSpec(panierloadbalancer.FieldID, field.TypeInt))
	if ps := plbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plbu.mutation.Quantity(); ok {
		_spec.SetField(panierloadbalancer.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := plbu.mutation.AddedQuantity(); ok {
		_spec.AddField(panierloadbalancer.FieldQuantity, field.TypeInt, value)
	}
	if plbu.mutation.PanierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.PanierTable,
			Columns: []string{panierloadbalancer.PanierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panier.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plbu.mutation.PanierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.PanierTable,
			Columns: []string{panierloadbalancer.PanierColumn},
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
	if plbu.mutation.LoadBalancerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.LoadBalancerTable,
			Columns: []string{panierloadbalancer.LoadBalancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plbu.mutation.LoadBalancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.LoadBalancerTable,
			Columns: []string{panierloadbalancer.LoadBalancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, plbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{panierloadbalancer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	plbu.mutation.done = true
	return n, nil
}

// PanierLoadBalancerUpdateOne is the builder for updating a single PanierLoadBalancer entity.
type PanierLoadBalancerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PanierLoadBalancerMutation
}

// SetQuantity sets the "quantity" field.
func (plbuo *PanierLoadBalancerUpdateOne) SetQuantity(i int) *PanierLoadBalancerUpdateOne {
	plbuo.mutation.ResetQuantity()
	plbuo.mutation.SetQuantity(i)
	return plbuo
}

// AddQuantity adds i to the "quantity" field.
func (plbuo *PanierLoadBalancerUpdateOne) AddQuantity(i int) *PanierLoadBalancerUpdateOne {
	plbuo.mutation.AddQuantity(i)
	return plbuo
}

// SetPanierID sets the "panier" edge to the Panier entity by ID.
func (plbuo *PanierLoadBalancerUpdateOne) SetPanierID(id int) *PanierLoadBalancerUpdateOne {
	plbuo.mutation.SetPanierID(id)
	return plbuo
}

// SetPanier sets the "panier" edge to the Panier entity.
func (plbuo *PanierLoadBalancerUpdateOne) SetPanier(p *Panier) *PanierLoadBalancerUpdateOne {
	return plbuo.SetPanierID(p.ID)
}

// SetLoadBalancerID sets the "loadBalancer" edge to the LoadBalancer entity by ID.
func (plbuo *PanierLoadBalancerUpdateOne) SetLoadBalancerID(id int) *PanierLoadBalancerUpdateOne {
	plbuo.mutation.SetLoadBalancerID(id)
	return plbuo
}

// SetLoadBalancer sets the "loadBalancer" edge to the LoadBalancer entity.
func (plbuo *PanierLoadBalancerUpdateOne) SetLoadBalancer(l *LoadBalancer) *PanierLoadBalancerUpdateOne {
	return plbuo.SetLoadBalancerID(l.ID)
}

// Mutation returns the PanierLoadBalancerMutation object of the builder.
func (plbuo *PanierLoadBalancerUpdateOne) Mutation() *PanierLoadBalancerMutation {
	return plbuo.mutation
}

// ClearPanier clears the "panier" edge to the Panier entity.
func (plbuo *PanierLoadBalancerUpdateOne) ClearPanier() *PanierLoadBalancerUpdateOne {
	plbuo.mutation.ClearPanier()
	return plbuo
}

// ClearLoadBalancer clears the "loadBalancer" edge to the LoadBalancer entity.
func (plbuo *PanierLoadBalancerUpdateOne) ClearLoadBalancer() *PanierLoadBalancerUpdateOne {
	plbuo.mutation.ClearLoadBalancer()
	return plbuo
}

// Where appends a list predicates to the PanierLoadBalancerUpdate builder.
func (plbuo *PanierLoadBalancerUpdateOne) Where(ps ...predicate.PanierLoadBalancer) *PanierLoadBalancerUpdateOne {
	plbuo.mutation.Where(ps...)
	return plbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (plbuo *PanierLoadBalancerUpdateOne) Select(field string, fields ...string) *PanierLoadBalancerUpdateOne {
	plbuo.fields = append([]string{field}, fields...)
	return plbuo
}

// Save executes the query and returns the updated PanierLoadBalancer entity.
func (plbuo *PanierLoadBalancerUpdateOne) Save(ctx context.Context) (*PanierLoadBalancer, error) {
	return withHooks(ctx, plbuo.sqlSave, plbuo.mutation, plbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (plbuo *PanierLoadBalancerUpdateOne) SaveX(ctx context.Context) *PanierLoadBalancer {
	node, err := plbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (plbuo *PanierLoadBalancerUpdateOne) Exec(ctx context.Context) error {
	_, err := plbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plbuo *PanierLoadBalancerUpdateOne) ExecX(ctx context.Context) {
	if err := plbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plbuo *PanierLoadBalancerUpdateOne) check() error {
	if _, ok := plbuo.mutation.PanierID(); plbuo.mutation.PanierCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierLoadBalancer.panier"`)
	}
	if _, ok := plbuo.mutation.LoadBalancerID(); plbuo.mutation.LoadBalancerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PanierLoadBalancer.loadBalancer"`)
	}
	return nil
}

func (plbuo *PanierLoadBalancerUpdateOne) sqlSave(ctx context.Context) (_node *PanierLoadBalancer, err error) {
	if err := plbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(panierloadbalancer.Table, panierloadbalancer.Columns, sqlgraph.NewFieldSpec(panierloadbalancer.FieldID, field.TypeInt))
	id, ok := plbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PanierLoadBalancer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := plbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, panierloadbalancer.FieldID)
		for _, f := range fields {
			if !panierloadbalancer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != panierloadbalancer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := plbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plbuo.mutation.Quantity(); ok {
		_spec.SetField(panierloadbalancer.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := plbuo.mutation.AddedQuantity(); ok {
		_spec.AddField(panierloadbalancer.FieldQuantity, field.TypeInt, value)
	}
	if plbuo.mutation.PanierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.PanierTable,
			Columns: []string{panierloadbalancer.PanierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panier.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plbuo.mutation.PanierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.PanierTable,
			Columns: []string{panierloadbalancer.PanierColumn},
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
	if plbuo.mutation.LoadBalancerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.LoadBalancerTable,
			Columns: []string{panierloadbalancer.LoadBalancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plbuo.mutation.LoadBalancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   panierloadbalancer.LoadBalancerTable,
			Columns: []string{panierloadbalancer.LoadBalancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PanierLoadBalancer{config: plbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, plbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{panierloadbalancer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	plbuo.mutation.done = true
	return _node, nil
}
