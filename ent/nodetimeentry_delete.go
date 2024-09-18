// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/realm76/ranger/ent/nodetimeentry"
	"github.com/realm76/ranger/ent/predicate"
)

// NodeTimeEntryDelete is the builder for deleting a NodeTimeEntry entity.
type NodeTimeEntryDelete struct {
	config
	hooks    []Hook
	mutation *NodeTimeEntryMutation
}

// Where appends a list predicates to the NodeTimeEntryDelete builder.
func (nted *NodeTimeEntryDelete) Where(ps ...predicate.NodeTimeEntry) *NodeTimeEntryDelete {
	nted.mutation.Where(ps...)
	return nted
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nted *NodeTimeEntryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nted.sqlExec, nted.mutation, nted.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nted *NodeTimeEntryDelete) ExecX(ctx context.Context) int {
	n, err := nted.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nted *NodeTimeEntryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(nodetimeentry.Table, sqlgraph.NewFieldSpec(nodetimeentry.FieldID, field.TypeInt))
	if ps := nted.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nted.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nted.mutation.done = true
	return affected, err
}

// NodeTimeEntryDeleteOne is the builder for deleting a single NodeTimeEntry entity.
type NodeTimeEntryDeleteOne struct {
	nted *NodeTimeEntryDelete
}

// Where appends a list predicates to the NodeTimeEntryDelete builder.
func (ntedo *NodeTimeEntryDeleteOne) Where(ps ...predicate.NodeTimeEntry) *NodeTimeEntryDeleteOne {
	ntedo.nted.mutation.Where(ps...)
	return ntedo
}

// Exec executes the deletion query.
func (ntedo *NodeTimeEntryDeleteOne) Exec(ctx context.Context) error {
	n, err := ntedo.nted.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nodetimeentry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ntedo *NodeTimeEntryDeleteOne) ExecX(ctx context.Context) {
	if err := ntedo.Exec(ctx); err != nil {
		panic(err)
	}
}
