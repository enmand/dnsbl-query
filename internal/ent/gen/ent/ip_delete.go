// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/enmand/dnsbl-query/internel/ent/gen/ent/ip"
	"github.com/enmand/dnsbl-query/internel/ent/gen/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// IPDelete is the builder for deleting a IP entity.
type IPDelete struct {
	config
	hooks    []Hook
	mutation *IPMutation
}

// Where adds a new predicate to the delete builder.
func (id *IPDelete) Where(ps ...predicate.IP) *IPDelete {
	id.mutation.predicates = append(id.mutation.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IPDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(id.hooks) - 1; i >= 0; i-- {
			mut = id.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IPDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IPDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ip.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: ip.FieldID,
			},
		},
	}
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, id.driver, _spec)
}

// IPDeleteOne is the builder for deleting a single IP entity.
type IPDeleteOne struct {
	id *IPDelete
}

// Exec executes the deletion query.
func (ido *IPDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ip.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IPDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
