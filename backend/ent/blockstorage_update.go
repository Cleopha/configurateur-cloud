// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/blockstorage"
	"backend/ent/panierblockstorage"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlockStorageUpdate is the builder for updating BlockStorage entities.
type BlockStorageUpdate struct {
	config
	hooks    []Hook
	mutation *BlockStorageMutation
}

// Where appends a list predicates to the BlockStorageUpdate builder.
func (bsu *BlockStorageUpdate) Where(ps ...predicate.BlockStorage) *BlockStorageUpdate {
	bsu.mutation.Where(ps...)
	return bsu
}

// SetName sets the "Name" field.
func (bsu *BlockStorageUpdate) SetName(s string) *BlockStorageUpdate {
	bsu.mutation.SetName(s)
	return bsu
}

// SetIOPS sets the "IOPS" field.
func (bsu *BlockStorageUpdate) SetIOPS(i int) *BlockStorageUpdate {
	bsu.mutation.ResetIOPS()
	bsu.mutation.SetIOPS(i)
	return bsu
}

// AddIOPS adds i to the "IOPS" field.
func (bsu *BlockStorageUpdate) AddIOPS(i int) *BlockStorageUpdate {
	bsu.mutation.AddIOPS(i)
	return bsu
}

// SetBandwidth sets the "Bandwidth" field.
func (bsu *BlockStorageUpdate) SetBandwidth(f float64) *BlockStorageUpdate {
	bsu.mutation.ResetBandwidth()
	bsu.mutation.SetBandwidth(f)
	return bsu
}

// AddBandwidth adds f to the "Bandwidth" field.
func (bsu *BlockStorageUpdate) AddBandwidth(f float64) *BlockStorageUpdate {
	bsu.mutation.AddBandwidth(f)
	return bsu
}

// SetPrice sets the "Price" field.
func (bsu *BlockStorageUpdate) SetPrice(f float64) *BlockStorageUpdate {
	bsu.mutation.ResetPrice()
	bsu.mutation.SetPrice(f)
	return bsu
}

// AddPrice adds f to the "Price" field.
func (bsu *BlockStorageUpdate) AddPrice(f float64) *BlockStorageUpdate {
	bsu.mutation.AddPrice(f)
	return bsu
}

// AddPanierBlockStorageIDs adds the "panierBlockStorage" edge to the PanierBlockStorage entity by IDs.
func (bsu *BlockStorageUpdate) AddPanierBlockStorageIDs(ids ...int) *BlockStorageUpdate {
	bsu.mutation.AddPanierBlockStorageIDs(ids...)
	return bsu
}

// AddPanierBlockStorage adds the "panierBlockStorage" edges to the PanierBlockStorage entity.
func (bsu *BlockStorageUpdate) AddPanierBlockStorage(p ...*PanierBlockStorage) *BlockStorageUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bsu.AddPanierBlockStorageIDs(ids...)
}

// Mutation returns the BlockStorageMutation object of the builder.
func (bsu *BlockStorageUpdate) Mutation() *BlockStorageMutation {
	return bsu.mutation
}

// ClearPanierBlockStorage clears all "panierBlockStorage" edges to the PanierBlockStorage entity.
func (bsu *BlockStorageUpdate) ClearPanierBlockStorage() *BlockStorageUpdate {
	bsu.mutation.ClearPanierBlockStorage()
	return bsu
}

// RemovePanierBlockStorageIDs removes the "panierBlockStorage" edge to PanierBlockStorage entities by IDs.
func (bsu *BlockStorageUpdate) RemovePanierBlockStorageIDs(ids ...int) *BlockStorageUpdate {
	bsu.mutation.RemovePanierBlockStorageIDs(ids...)
	return bsu
}

// RemovePanierBlockStorage removes "panierBlockStorage" edges to PanierBlockStorage entities.
func (bsu *BlockStorageUpdate) RemovePanierBlockStorage(p ...*PanierBlockStorage) *BlockStorageUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bsu.RemovePanierBlockStorageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bsu *BlockStorageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bsu.sqlSave, bsu.mutation, bsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bsu *BlockStorageUpdate) SaveX(ctx context.Context) int {
	affected, err := bsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bsu *BlockStorageUpdate) Exec(ctx context.Context) error {
	_, err := bsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsu *BlockStorageUpdate) ExecX(ctx context.Context) {
	if err := bsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsu *BlockStorageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blockstorage.Table, blockstorage.Columns, sqlgraph.NewFieldSpec(blockstorage.FieldID, field.TypeInt))
	if ps := bsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bsu.mutation.Name(); ok {
		_spec.SetField(blockstorage.FieldName, field.TypeString, value)
	}
	if value, ok := bsu.mutation.IOPS(); ok {
		_spec.SetField(blockstorage.FieldIOPS, field.TypeInt, value)
	}
	if value, ok := bsu.mutation.AddedIOPS(); ok {
		_spec.AddField(blockstorage.FieldIOPS, field.TypeInt, value)
	}
	if value, ok := bsu.mutation.Bandwidth(); ok {
		_spec.SetField(blockstorage.FieldBandwidth, field.TypeFloat64, value)
	}
	if value, ok := bsu.mutation.AddedBandwidth(); ok {
		_spec.AddField(blockstorage.FieldBandwidth, field.TypeFloat64, value)
	}
	if value, ok := bsu.mutation.Price(); ok {
		_spec.SetField(blockstorage.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := bsu.mutation.AddedPrice(); ok {
		_spec.AddField(blockstorage.FieldPrice, field.TypeFloat64, value)
	}
	if bsu.mutation.PanierBlockStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockstorage.PanierBlockStorageTable,
			Columns: []string{blockstorage.PanierBlockStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsu.mutation.RemovedPanierBlockStorageIDs(); len(nodes) > 0 && !bsu.mutation.PanierBlockStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockstorage.PanierBlockStorageTable,
			Columns: []string{blockstorage.PanierBlockStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsu.mutation.PanierBlockStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockstorage.PanierBlockStorageTable,
			Columns: []string{blockstorage.PanierBlockStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blockstorage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bsu.mutation.done = true
	return n, nil
}

// BlockStorageUpdateOne is the builder for updating a single BlockStorage entity.
type BlockStorageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlockStorageMutation
}

// SetName sets the "Name" field.
func (bsuo *BlockStorageUpdateOne) SetName(s string) *BlockStorageUpdateOne {
	bsuo.mutation.SetName(s)
	return bsuo
}

// SetIOPS sets the "IOPS" field.
func (bsuo *BlockStorageUpdateOne) SetIOPS(i int) *BlockStorageUpdateOne {
	bsuo.mutation.ResetIOPS()
	bsuo.mutation.SetIOPS(i)
	return bsuo
}

// AddIOPS adds i to the "IOPS" field.
func (bsuo *BlockStorageUpdateOne) AddIOPS(i int) *BlockStorageUpdateOne {
	bsuo.mutation.AddIOPS(i)
	return bsuo
}

// SetBandwidth sets the "Bandwidth" field.
func (bsuo *BlockStorageUpdateOne) SetBandwidth(f float64) *BlockStorageUpdateOne {
	bsuo.mutation.ResetBandwidth()
	bsuo.mutation.SetBandwidth(f)
	return bsuo
}

// AddBandwidth adds f to the "Bandwidth" field.
func (bsuo *BlockStorageUpdateOne) AddBandwidth(f float64) *BlockStorageUpdateOne {
	bsuo.mutation.AddBandwidth(f)
	return bsuo
}

// SetPrice sets the "Price" field.
func (bsuo *BlockStorageUpdateOne) SetPrice(f float64) *BlockStorageUpdateOne {
	bsuo.mutation.ResetPrice()
	bsuo.mutation.SetPrice(f)
	return bsuo
}

// AddPrice adds f to the "Price" field.
func (bsuo *BlockStorageUpdateOne) AddPrice(f float64) *BlockStorageUpdateOne {
	bsuo.mutation.AddPrice(f)
	return bsuo
}

// AddPanierBlockStorageIDs adds the "panierBlockStorage" edge to the PanierBlockStorage entity by IDs.
func (bsuo *BlockStorageUpdateOne) AddPanierBlockStorageIDs(ids ...int) *BlockStorageUpdateOne {
	bsuo.mutation.AddPanierBlockStorageIDs(ids...)
	return bsuo
}

// AddPanierBlockStorage adds the "panierBlockStorage" edges to the PanierBlockStorage entity.
func (bsuo *BlockStorageUpdateOne) AddPanierBlockStorage(p ...*PanierBlockStorage) *BlockStorageUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bsuo.AddPanierBlockStorageIDs(ids...)
}

// Mutation returns the BlockStorageMutation object of the builder.
func (bsuo *BlockStorageUpdateOne) Mutation() *BlockStorageMutation {
	return bsuo.mutation
}

// ClearPanierBlockStorage clears all "panierBlockStorage" edges to the PanierBlockStorage entity.
func (bsuo *BlockStorageUpdateOne) ClearPanierBlockStorage() *BlockStorageUpdateOne {
	bsuo.mutation.ClearPanierBlockStorage()
	return bsuo
}

// RemovePanierBlockStorageIDs removes the "panierBlockStorage" edge to PanierBlockStorage entities by IDs.
func (bsuo *BlockStorageUpdateOne) RemovePanierBlockStorageIDs(ids ...int) *BlockStorageUpdateOne {
	bsuo.mutation.RemovePanierBlockStorageIDs(ids...)
	return bsuo
}

// RemovePanierBlockStorage removes "panierBlockStorage" edges to PanierBlockStorage entities.
func (bsuo *BlockStorageUpdateOne) RemovePanierBlockStorage(p ...*PanierBlockStorage) *BlockStorageUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bsuo.RemovePanierBlockStorageIDs(ids...)
}

// Where appends a list predicates to the BlockStorageUpdate builder.
func (bsuo *BlockStorageUpdateOne) Where(ps ...predicate.BlockStorage) *BlockStorageUpdateOne {
	bsuo.mutation.Where(ps...)
	return bsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bsuo *BlockStorageUpdateOne) Select(field string, fields ...string) *BlockStorageUpdateOne {
	bsuo.fields = append([]string{field}, fields...)
	return bsuo
}

// Save executes the query and returns the updated BlockStorage entity.
func (bsuo *BlockStorageUpdateOne) Save(ctx context.Context) (*BlockStorage, error) {
	return withHooks(ctx, bsuo.sqlSave, bsuo.mutation, bsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bsuo *BlockStorageUpdateOne) SaveX(ctx context.Context) *BlockStorage {
	node, err := bsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bsuo *BlockStorageUpdateOne) Exec(ctx context.Context) error {
	_, err := bsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsuo *BlockStorageUpdateOne) ExecX(ctx context.Context) {
	if err := bsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsuo *BlockStorageUpdateOne) sqlSave(ctx context.Context) (_node *BlockStorage, err error) {
	_spec := sqlgraph.NewUpdateSpec(blockstorage.Table, blockstorage.Columns, sqlgraph.NewFieldSpec(blockstorage.FieldID, field.TypeInt))
	id, ok := bsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BlockStorage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blockstorage.FieldID)
		for _, f := range fields {
			if !blockstorage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blockstorage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bsuo.mutation.Name(); ok {
		_spec.SetField(blockstorage.FieldName, field.TypeString, value)
	}
	if value, ok := bsuo.mutation.IOPS(); ok {
		_spec.SetField(blockstorage.FieldIOPS, field.TypeInt, value)
	}
	if value, ok := bsuo.mutation.AddedIOPS(); ok {
		_spec.AddField(blockstorage.FieldIOPS, field.TypeInt, value)
	}
	if value, ok := bsuo.mutation.Bandwidth(); ok {
		_spec.SetField(blockstorage.FieldBandwidth, field.TypeFloat64, value)
	}
	if value, ok := bsuo.mutation.AddedBandwidth(); ok {
		_spec.AddField(blockstorage.FieldBandwidth, field.TypeFloat64, value)
	}
	if value, ok := bsuo.mutation.Price(); ok {
		_spec.SetField(blockstorage.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := bsuo.mutation.AddedPrice(); ok {
		_spec.AddField(blockstorage.FieldPrice, field.TypeFloat64, value)
	}
	if bsuo.mutation.PanierBlockStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockstorage.PanierBlockStorageTable,
			Columns: []string{blockstorage.PanierBlockStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsuo.mutation.RemovedPanierBlockStorageIDs(); len(nodes) > 0 && !bsuo.mutation.PanierBlockStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockstorage.PanierBlockStorageTable,
			Columns: []string{blockstorage.PanierBlockStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsuo.mutation.PanierBlockStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockstorage.PanierBlockStorageTable,
			Columns: []string{blockstorage.PanierBlockStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BlockStorage{config: bsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blockstorage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bsuo.mutation.done = true
	return _node, nil
}
