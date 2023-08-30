// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/loadbalancer"
	"backend/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoadBalancerDelete is the builder for deleting a LoadBalancer entity.
type LoadBalancerDelete struct {
	config
	hooks    []Hook
	mutation *LoadBalancerMutation
}

// Where appends a list predicates to the LoadBalancerDelete builder.
func (lbd *LoadBalancerDelete) Where(ps ...predicate.LoadBalancer) *LoadBalancerDelete {
	lbd.mutation.Where(ps...)
	return lbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lbd *LoadBalancerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lbd.sqlExec, lbd.mutation, lbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lbd *LoadBalancerDelete) ExecX(ctx context.Context) int {
	n, err := lbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lbd *LoadBalancerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(loadbalancer.Table, sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeInt))
	if ps := lbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lbd.mutation.done = true
	return affected, err
}

// LoadBalancerDeleteOne is the builder for deleting a single LoadBalancer entity.
type LoadBalancerDeleteOne struct {
	lbd *LoadBalancerDelete
}

// Where appends a list predicates to the LoadBalancerDelete builder.
func (lbdo *LoadBalancerDeleteOne) Where(ps ...predicate.LoadBalancer) *LoadBalancerDeleteOne {
	lbdo.lbd.mutation.Where(ps...)
	return lbdo
}

// Exec executes the deletion query.
func (lbdo *LoadBalancerDeleteOne) Exec(ctx context.Context) error {
	n, err := lbdo.lbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loadbalancer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lbdo *LoadBalancerDeleteOne) ExecX(ctx context.Context) {
	if err := lbdo.Exec(ctx); err != nil {
		panic(err)
	}
}