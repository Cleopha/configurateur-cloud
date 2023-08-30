// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/panierobjectstorage"
	"backend/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierObjectStorageDelete is the builder for deleting a PanierObjectStorage entity.
type PanierObjectStorageDelete struct {
	config
	hooks    []Hook
	mutation *PanierObjectStorageMutation
}

// Where appends a list predicates to the PanierObjectStorageDelete builder.
func (posd *PanierObjectStorageDelete) Where(ps ...predicate.PanierObjectStorage) *PanierObjectStorageDelete {
	posd.mutation.Where(ps...)
	return posd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (posd *PanierObjectStorageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, posd.sqlExec, posd.mutation, posd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (posd *PanierObjectStorageDelete) ExecX(ctx context.Context) int {
	n, err := posd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (posd *PanierObjectStorageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(panierobjectstorage.Table, sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt))
	if ps := posd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, posd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	posd.mutation.done = true
	return affected, err
}

// PanierObjectStorageDeleteOne is the builder for deleting a single PanierObjectStorage entity.
type PanierObjectStorageDeleteOne struct {
	posd *PanierObjectStorageDelete
}

// Where appends a list predicates to the PanierObjectStorageDelete builder.
func (posdo *PanierObjectStorageDeleteOne) Where(ps ...predicate.PanierObjectStorage) *PanierObjectStorageDeleteOne {
	posdo.posd.mutation.Where(ps...)
	return posdo
}

// Exec executes the deletion query.
func (posdo *PanierObjectStorageDeleteOne) Exec(ctx context.Context) error {
	n, err := posdo.posd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{panierobjectstorage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (posdo *PanierObjectStorageDeleteOne) ExecX(ctx context.Context) {
	if err := posdo.Exec(ctx); err != nil {
		panic(err)
	}
}
