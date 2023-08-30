// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/blockstorage"
	"backend/ent/panierblockstorage"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlockStorageCreate is the builder for creating a BlockStorage entity.
type BlockStorageCreate struct {
	config
	mutation *BlockStorageMutation
	hooks    []Hook
}

// SetName sets the "Name" field.
func (bsc *BlockStorageCreate) SetName(s string) *BlockStorageCreate {
	bsc.mutation.SetName(s)
	return bsc
}

// SetIOPS sets the "IOPS" field.
func (bsc *BlockStorageCreate) SetIOPS(i int) *BlockStorageCreate {
	bsc.mutation.SetIOPS(i)
	return bsc
}

// SetBandwidth sets the "Bandwidth" field.
func (bsc *BlockStorageCreate) SetBandwidth(f float64) *BlockStorageCreate {
	bsc.mutation.SetBandwidth(f)
	return bsc
}

// SetPrice sets the "Price" field.
func (bsc *BlockStorageCreate) SetPrice(f float64) *BlockStorageCreate {
	bsc.mutation.SetPrice(f)
	return bsc
}

// SetID sets the "id" field.
func (bsc *BlockStorageCreate) SetID(i int) *BlockStorageCreate {
	bsc.mutation.SetID(i)
	return bsc
}

// AddPanierBlockStorageIDs adds the "panierBlockStorage" edge to the PanierBlockStorage entity by IDs.
func (bsc *BlockStorageCreate) AddPanierBlockStorageIDs(ids ...int) *BlockStorageCreate {
	bsc.mutation.AddPanierBlockStorageIDs(ids...)
	return bsc
}

// AddPanierBlockStorage adds the "panierBlockStorage" edges to the PanierBlockStorage entity.
func (bsc *BlockStorageCreate) AddPanierBlockStorage(p ...*PanierBlockStorage) *BlockStorageCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bsc.AddPanierBlockStorageIDs(ids...)
}

// Mutation returns the BlockStorageMutation object of the builder.
func (bsc *BlockStorageCreate) Mutation() *BlockStorageMutation {
	return bsc.mutation
}

// Save creates the BlockStorage in the database.
func (bsc *BlockStorageCreate) Save(ctx context.Context) (*BlockStorage, error) {
	return withHooks(ctx, bsc.sqlSave, bsc.mutation, bsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bsc *BlockStorageCreate) SaveX(ctx context.Context) *BlockStorage {
	v, err := bsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bsc *BlockStorageCreate) Exec(ctx context.Context) error {
	_, err := bsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsc *BlockStorageCreate) ExecX(ctx context.Context) {
	if err := bsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bsc *BlockStorageCreate) check() error {
	if _, ok := bsc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "BlockStorage.Name"`)}
	}
	if _, ok := bsc.mutation.IOPS(); !ok {
		return &ValidationError{Name: "IOPS", err: errors.New(`ent: missing required field "BlockStorage.IOPS"`)}
	}
	if _, ok := bsc.mutation.Bandwidth(); !ok {
		return &ValidationError{Name: "Bandwidth", err: errors.New(`ent: missing required field "BlockStorage.Bandwidth"`)}
	}
	if _, ok := bsc.mutation.Price(); !ok {
		return &ValidationError{Name: "Price", err: errors.New(`ent: missing required field "BlockStorage.Price"`)}
	}
	return nil
}

func (bsc *BlockStorageCreate) sqlSave(ctx context.Context) (*BlockStorage, error) {
	if err := bsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	bsc.mutation.id = &_node.ID
	bsc.mutation.done = true
	return _node, nil
}

func (bsc *BlockStorageCreate) createSpec() (*BlockStorage, *sqlgraph.CreateSpec) {
	var (
		_node = &BlockStorage{config: bsc.config}
		_spec = sqlgraph.NewCreateSpec(blockstorage.Table, sqlgraph.NewFieldSpec(blockstorage.FieldID, field.TypeInt))
	)
	if id, ok := bsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bsc.mutation.Name(); ok {
		_spec.SetField(blockstorage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bsc.mutation.IOPS(); ok {
		_spec.SetField(blockstorage.FieldIOPS, field.TypeInt, value)
		_node.IOPS = value
	}
	if value, ok := bsc.mutation.Bandwidth(); ok {
		_spec.SetField(blockstorage.FieldBandwidth, field.TypeFloat64, value)
		_node.Bandwidth = value
	}
	if value, ok := bsc.mutation.Price(); ok {
		_spec.SetField(blockstorage.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if nodes := bsc.mutation.PanierBlockStorageIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlockStorageCreateBulk is the builder for creating many BlockStorage entities in bulk.
type BlockStorageCreateBulk struct {
	config
	builders []*BlockStorageCreate
}

// Save creates the BlockStorage entities in the database.
func (bscb *BlockStorageCreateBulk) Save(ctx context.Context) ([]*BlockStorage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bscb.builders))
	nodes := make([]*BlockStorage, len(bscb.builders))
	mutators := make([]Mutator, len(bscb.builders))
	for i := range bscb.builders {
		func(i int, root context.Context) {
			builder := bscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockStorageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bscb *BlockStorageCreateBulk) SaveX(ctx context.Context) []*BlockStorage {
	v, err := bscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bscb *BlockStorageCreateBulk) Exec(ctx context.Context) error {
	_, err := bscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bscb *BlockStorageCreateBulk) ExecX(ctx context.Context) {
	if err := bscb.Exec(ctx); err != nil {
		panic(err)
	}
}
