// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/operation"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// OperationCreate is the builder for creating a Operation entity.
type OperationCreate struct {
	config
	mutation *OperationMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (oc *OperationCreate) SetCreatedAt(t time.Time) *OperationCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (oc *OperationCreate) SetNillableCreatedAt(t *time.Time) *OperationCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the updated_at field.
func (oc *OperationCreate) SetUpdatedAt(t time.Time) *OperationCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (oc *OperationCreate) SetNillableUpdatedAt(t *time.Time) *OperationCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetType sets the type field.
func (oc *OperationCreate) SetType(o operation.Type) *OperationCreate {
	oc.mutation.SetType(o)
	return oc
}

// SetIPAddress sets the ip_address field.
func (oc *OperationCreate) SetIPAddress(s string) *OperationCreate {
	oc.mutation.SetIPAddress(s)
	return oc
}

// SetNillableIPAddress sets the ip_address field if the given value is not nil.
func (oc *OperationCreate) SetNillableIPAddress(s *string) *OperationCreate {
	if s != nil {
		oc.SetIPAddress(*s)
	}
	return oc
}

// SetStatus sets the status field.
func (oc *OperationCreate) SetStatus(o operation.Status) *OperationCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetError sets the error field.
func (oc *OperationCreate) SetError(s string) *OperationCreate {
	oc.mutation.SetError(s)
	return oc
}

// SetNillableError sets the error field if the given value is not nil.
func (oc *OperationCreate) SetNillableError(s *string) *OperationCreate {
	if s != nil {
		oc.SetError(*s)
	}
	return oc
}

// SetDoneAt sets the done_at field.
func (oc *OperationCreate) SetDoneAt(t time.Time) *OperationCreate {
	oc.mutation.SetDoneAt(t)
	return oc
}

// SetNillableDoneAt sets the done_at field if the given value is not nil.
func (oc *OperationCreate) SetNillableDoneAt(t *time.Time) *OperationCreate {
	if t != nil {
		oc.SetDoneAt(*t)
	}
	return oc
}

// SetID sets the id field.
func (oc *OperationCreate) SetID(u uuid.UUID) *OperationCreate {
	oc.mutation.SetID(u)
	return oc
}

// Mutation returns the OperationMutation object of the builder.
func (oc *OperationCreate) Mutation() *OperationMutation {
	return oc.mutation
}

// Save creates the Operation in the database.
func (oc *OperationCreate) Save(ctx context.Context) (*Operation, error) {
	var (
		err  error
		node *Operation
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OperationCreate) SaveX(ctx context.Context) *Operation {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (oc *OperationCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := operation.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := operation.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		v := operation.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OperationCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if v, ok := oc.mutation.GetType(); ok {
		if err := operation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := operation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (oc *OperationCreate) sqlSave(ctx context.Context) (*Operation, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (oc *OperationCreate) createSpec() (*Operation, *sqlgraph.CreateSpec) {
	var (
		_node = &Operation{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: operation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: operation.FieldID,
			},
		}
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: operation.FieldType,
		})
		_node.Type = value
	}
	if value, ok := oc.mutation.IPAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operation.FieldIPAddress,
		})
		_node.IPAddress = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: operation.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := oc.mutation.Error(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operation.FieldError,
		})
		_node.Error = value
	}
	if value, ok := oc.mutation.DoneAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldDoneAt,
		})
		_node.DoneAt = value
	}
	return _node, _spec
}

// OperationCreateBulk is the builder for creating a bulk of Operation entities.
type OperationCreateBulk struct {
	config
	builders []*OperationCreate
}

// Save creates the Operation entities in the database.
func (ocb *OperationCreateBulk) Save(ctx context.Context) ([]*Operation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Operation, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ocb *OperationCreateBulk) SaveX(ctx context.Context) []*Operation {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
