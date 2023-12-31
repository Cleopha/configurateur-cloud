// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/objectstorage"
	"backend/ent/panierobjectstorage"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ObjectStorageUpdate is the builder for updating ObjectStorage entities.
type ObjectStorageUpdate struct {
	config
	hooks    []Hook
	mutation *ObjectStorageMutation
}

// Where appends a list predicates to the ObjectStorageUpdate builder.
func (osu *ObjectStorageUpdate) Where(ps ...predicate.ObjectStorage) *ObjectStorageUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetName sets the "Name" field.
func (osu *ObjectStorageUpdate) SetName(s string) *ObjectStorageUpdate {
	osu.mutation.SetName(s)
	return osu
}

// SetPrice sets the "Price" field.
func (osu *ObjectStorageUpdate) SetPrice(f float64) *ObjectStorageUpdate {
	osu.mutation.ResetPrice()
	osu.mutation.SetPrice(f)
	return osu
}

// AddPrice adds f to the "Price" field.
func (osu *ObjectStorageUpdate) AddPrice(f float64) *ObjectStorageUpdate {
	osu.mutation.AddPrice(f)
	return osu
}

// AddPanierObjectStorageIDs adds the "panierObjectStorage" edge to the PanierObjectStorage entity by IDs.
func (osu *ObjectStorageUpdate) AddPanierObjectStorageIDs(ids ...int) *ObjectStorageUpdate {
	osu.mutation.AddPanierObjectStorageIDs(ids...)
	return osu
}

// AddPanierObjectStorage adds the "panierObjectStorage" edges to the PanierObjectStorage entity.
func (osu *ObjectStorageUpdate) AddPanierObjectStorage(p ...*PanierObjectStorage) *ObjectStorageUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return osu.AddPanierObjectStorageIDs(ids...)
}

// Mutation returns the ObjectStorageMutation object of the builder.
func (osu *ObjectStorageUpdate) Mutation() *ObjectStorageMutation {
	return osu.mutation
}

// ClearPanierObjectStorage clears all "panierObjectStorage" edges to the PanierObjectStorage entity.
func (osu *ObjectStorageUpdate) ClearPanierObjectStorage() *ObjectStorageUpdate {
	osu.mutation.ClearPanierObjectStorage()
	return osu
}

// RemovePanierObjectStorageIDs removes the "panierObjectStorage" edge to PanierObjectStorage entities by IDs.
func (osu *ObjectStorageUpdate) RemovePanierObjectStorageIDs(ids ...int) *ObjectStorageUpdate {
	osu.mutation.RemovePanierObjectStorageIDs(ids...)
	return osu
}

// RemovePanierObjectStorage removes "panierObjectStorage" edges to PanierObjectStorage entities.
func (osu *ObjectStorageUpdate) RemovePanierObjectStorage(p ...*PanierObjectStorage) *ObjectStorageUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return osu.RemovePanierObjectStorageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *ObjectStorageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, osu.sqlSave, osu.mutation, osu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osu *ObjectStorageUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *ObjectStorageUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *ObjectStorageUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (osu *ObjectStorageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(objectstorage.Table, objectstorage.Columns, sqlgraph.NewFieldSpec(objectstorage.FieldID, field.TypeInt))
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.Name(); ok {
		_spec.SetField(objectstorage.FieldName, field.TypeString, value)
	}
	if value, ok := osu.mutation.Price(); ok {
		_spec.SetField(objectstorage.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := osu.mutation.AddedPrice(); ok {
		_spec.AddField(objectstorage.FieldPrice, field.TypeFloat64, value)
	}
	if osu.mutation.PanierObjectStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   objectstorage.PanierObjectStorageTable,
			Columns: []string{objectstorage.PanierObjectStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.RemovedPanierObjectStorageIDs(); len(nodes) > 0 && !osu.mutation.PanierObjectStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   objectstorage.PanierObjectStorageTable,
			Columns: []string{objectstorage.PanierObjectStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.PanierObjectStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   objectstorage.PanierObjectStorageTable,
			Columns: []string{objectstorage.PanierObjectStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{objectstorage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	osu.mutation.done = true
	return n, nil
}

// ObjectStorageUpdateOne is the builder for updating a single ObjectStorage entity.
type ObjectStorageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ObjectStorageMutation
}

// SetName sets the "Name" field.
func (osuo *ObjectStorageUpdateOne) SetName(s string) *ObjectStorageUpdateOne {
	osuo.mutation.SetName(s)
	return osuo
}

// SetPrice sets the "Price" field.
func (osuo *ObjectStorageUpdateOne) SetPrice(f float64) *ObjectStorageUpdateOne {
	osuo.mutation.ResetPrice()
	osuo.mutation.SetPrice(f)
	return osuo
}

// AddPrice adds f to the "Price" field.
func (osuo *ObjectStorageUpdateOne) AddPrice(f float64) *ObjectStorageUpdateOne {
	osuo.mutation.AddPrice(f)
	return osuo
}

// AddPanierObjectStorageIDs adds the "panierObjectStorage" edge to the PanierObjectStorage entity by IDs.
func (osuo *ObjectStorageUpdateOne) AddPanierObjectStorageIDs(ids ...int) *ObjectStorageUpdateOne {
	osuo.mutation.AddPanierObjectStorageIDs(ids...)
	return osuo
}

// AddPanierObjectStorage adds the "panierObjectStorage" edges to the PanierObjectStorage entity.
func (osuo *ObjectStorageUpdateOne) AddPanierObjectStorage(p ...*PanierObjectStorage) *ObjectStorageUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return osuo.AddPanierObjectStorageIDs(ids...)
}

// Mutation returns the ObjectStorageMutation object of the builder.
func (osuo *ObjectStorageUpdateOne) Mutation() *ObjectStorageMutation {
	return osuo.mutation
}

// ClearPanierObjectStorage clears all "panierObjectStorage" edges to the PanierObjectStorage entity.
func (osuo *ObjectStorageUpdateOne) ClearPanierObjectStorage() *ObjectStorageUpdateOne {
	osuo.mutation.ClearPanierObjectStorage()
	return osuo
}

// RemovePanierObjectStorageIDs removes the "panierObjectStorage" edge to PanierObjectStorage entities by IDs.
func (osuo *ObjectStorageUpdateOne) RemovePanierObjectStorageIDs(ids ...int) *ObjectStorageUpdateOne {
	osuo.mutation.RemovePanierObjectStorageIDs(ids...)
	return osuo
}

// RemovePanierObjectStorage removes "panierObjectStorage" edges to PanierObjectStorage entities.
func (osuo *ObjectStorageUpdateOne) RemovePanierObjectStorage(p ...*PanierObjectStorage) *ObjectStorageUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return osuo.RemovePanierObjectStorageIDs(ids...)
}

// Where appends a list predicates to the ObjectStorageUpdate builder.
func (osuo *ObjectStorageUpdateOne) Where(ps ...predicate.ObjectStorage) *ObjectStorageUpdateOne {
	osuo.mutation.Where(ps...)
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *ObjectStorageUpdateOne) Select(field string, fields ...string) *ObjectStorageUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated ObjectStorage entity.
func (osuo *ObjectStorageUpdateOne) Save(ctx context.Context) (*ObjectStorage, error) {
	return withHooks(ctx, osuo.sqlSave, osuo.mutation, osuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *ObjectStorageUpdateOne) SaveX(ctx context.Context) *ObjectStorage {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *ObjectStorageUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *ObjectStorageUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (osuo *ObjectStorageUpdateOne) sqlSave(ctx context.Context) (_node *ObjectStorage, err error) {
	_spec := sqlgraph.NewUpdateSpec(objectstorage.Table, objectstorage.Columns, sqlgraph.NewFieldSpec(objectstorage.FieldID, field.TypeInt))
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ObjectStorage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, objectstorage.FieldID)
		for _, f := range fields {
			if !objectstorage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != objectstorage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.Name(); ok {
		_spec.SetField(objectstorage.FieldName, field.TypeString, value)
	}
	if value, ok := osuo.mutation.Price(); ok {
		_spec.SetField(objectstorage.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := osuo.mutation.AddedPrice(); ok {
		_spec.AddField(objectstorage.FieldPrice, field.TypeFloat64, value)
	}
	if osuo.mutation.PanierObjectStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   objectstorage.PanierObjectStorageTable,
			Columns: []string{objectstorage.PanierObjectStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.RemovedPanierObjectStorageIDs(); len(nodes) > 0 && !osuo.mutation.PanierObjectStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   objectstorage.PanierObjectStorageTable,
			Columns: []string{objectstorage.PanierObjectStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.PanierObjectStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   objectstorage.PanierObjectStorageTable,
			Columns: []string{objectstorage.PanierObjectStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ObjectStorage{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{objectstorage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	osuo.mutation.done = true
	return _node, nil
}
