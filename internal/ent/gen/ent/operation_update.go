// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/operation"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// OperationUpdate is the builder for updating Operation entities.
type OperationUpdate struct {
	config
	hooks    []Hook
	mutation *OperationMutation
}

// Where adds a new predicate for the builder.
func (ou *OperationUpdate) Where(ps ...predicate.Operation) *OperationUpdate {
	ou.mutation.predicates = append(ou.mutation.predicates, ps...)
	return ou
}

// SetCreatedAt sets the created_at field.
func (ou *OperationUpdate) SetCreatedAt(t time.Time) *OperationUpdate {
	ou.mutation.SetCreatedAt(t)
	return ou
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ou *OperationUpdate) SetNillableCreatedAt(t *time.Time) *OperationUpdate {
	if t != nil {
		ou.SetCreatedAt(*t)
	}
	return ou
}

// SetUpdatedAt sets the updated_at field.
func (ou *OperationUpdate) SetUpdatedAt(t time.Time) *OperationUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetType sets the type field.
func (ou *OperationUpdate) SetType(o operation.Type) *OperationUpdate {
	ou.mutation.SetType(o)
	return ou
}

// SetIPAddress sets the ip_address field.
func (ou *OperationUpdate) SetIPAddress(s string) *OperationUpdate {
	ou.mutation.SetIPAddress(s)
	return ou
}

// SetNillableIPAddress sets the ip_address field if the given value is not nil.
func (ou *OperationUpdate) SetNillableIPAddress(s *string) *OperationUpdate {
	if s != nil {
		ou.SetIPAddress(*s)
	}
	return ou
}

// ClearIPAddress clears the value of ip_address.
func (ou *OperationUpdate) ClearIPAddress() *OperationUpdate {
	ou.mutation.ClearIPAddress()
	return ou
}

// SetStatus sets the status field.
func (ou *OperationUpdate) SetStatus(o operation.Status) *OperationUpdate {
	ou.mutation.SetStatus(o)
	return ou
}

// SetDoneAt sets the done_at field.
func (ou *OperationUpdate) SetDoneAt(t time.Time) *OperationUpdate {
	ou.mutation.SetDoneAt(t)
	return ou
}

// SetNillableDoneAt sets the done_at field if the given value is not nil.
func (ou *OperationUpdate) SetNillableDoneAt(t *time.Time) *OperationUpdate {
	if t != nil {
		ou.SetDoneAt(*t)
	}
	return ou
}

// ClearDoneAt clears the value of done_at.
func (ou *OperationUpdate) ClearDoneAt() *OperationUpdate {
	ou.mutation.ClearDoneAt()
	return ou
}

// Mutation returns the OperationMutation object of the builder.
func (ou *OperationUpdate) Mutation() *OperationMutation {
	return ou.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OperationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OperationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OperationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OperationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OperationUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := operation.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OperationUpdate) check() error {
	if v, ok := ou.mutation.GetType(); ok {
		if err := operation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := ou.mutation.Status(); ok {
		if err := operation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ou *OperationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   operation.Table,
			Columns: operation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: operation.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: operation.FieldType,
		})
	}
	if value, ok := ou.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operation.FieldIPAddress,
		})
	}
	if ou.mutation.IPAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operation.FieldIPAddress,
		})
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: operation.FieldStatus,
		})
	}
	if value, ok := ou.mutation.DoneAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldDoneAt,
		})
	}
	if ou.mutation.DoneAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operation.FieldDoneAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{operation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OperationUpdateOne is the builder for updating a single Operation entity.
type OperationUpdateOne struct {
	config
	hooks    []Hook
	mutation *OperationMutation
}

// SetCreatedAt sets the created_at field.
func (ouo *OperationUpdateOne) SetCreatedAt(t time.Time) *OperationUpdateOne {
	ouo.mutation.SetCreatedAt(t)
	return ouo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ouo *OperationUpdateOne) SetNillableCreatedAt(t *time.Time) *OperationUpdateOne {
	if t != nil {
		ouo.SetCreatedAt(*t)
	}
	return ouo
}

// SetUpdatedAt sets the updated_at field.
func (ouo *OperationUpdateOne) SetUpdatedAt(t time.Time) *OperationUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetType sets the type field.
func (ouo *OperationUpdateOne) SetType(o operation.Type) *OperationUpdateOne {
	ouo.mutation.SetType(o)
	return ouo
}

// SetIPAddress sets the ip_address field.
func (ouo *OperationUpdateOne) SetIPAddress(s string) *OperationUpdateOne {
	ouo.mutation.SetIPAddress(s)
	return ouo
}

// SetNillableIPAddress sets the ip_address field if the given value is not nil.
func (ouo *OperationUpdateOne) SetNillableIPAddress(s *string) *OperationUpdateOne {
	if s != nil {
		ouo.SetIPAddress(*s)
	}
	return ouo
}

// ClearIPAddress clears the value of ip_address.
func (ouo *OperationUpdateOne) ClearIPAddress() *OperationUpdateOne {
	ouo.mutation.ClearIPAddress()
	return ouo
}

// SetStatus sets the status field.
func (ouo *OperationUpdateOne) SetStatus(o operation.Status) *OperationUpdateOne {
	ouo.mutation.SetStatus(o)
	return ouo
}

// SetDoneAt sets the done_at field.
func (ouo *OperationUpdateOne) SetDoneAt(t time.Time) *OperationUpdateOne {
	ouo.mutation.SetDoneAt(t)
	return ouo
}

// SetNillableDoneAt sets the done_at field if the given value is not nil.
func (ouo *OperationUpdateOne) SetNillableDoneAt(t *time.Time) *OperationUpdateOne {
	if t != nil {
		ouo.SetDoneAt(*t)
	}
	return ouo
}

// ClearDoneAt clears the value of done_at.
func (ouo *OperationUpdateOne) ClearDoneAt() *OperationUpdateOne {
	ouo.mutation.ClearDoneAt()
	return ouo
}

// Mutation returns the OperationMutation object of the builder.
func (ouo *OperationUpdateOne) Mutation() *OperationMutation {
	return ouo.mutation
}

// Save executes the query and returns the updated entity.
func (ouo *OperationUpdateOne) Save(ctx context.Context) (*Operation, error) {
	var (
		err  error
		node *Operation
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OperationUpdateOne) SaveX(ctx context.Context) *Operation {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OperationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OperationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OperationUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := operation.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OperationUpdateOne) check() error {
	if v, ok := ouo.mutation.GetType(); ok {
		if err := operation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := ouo.mutation.Status(); ok {
		if err := operation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ouo *OperationUpdateOne) sqlSave(ctx context.Context) (_node *Operation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   operation.Table,
			Columns: operation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: operation.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Operation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: operation.FieldType,
		})
	}
	if value, ok := ouo.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operation.FieldIPAddress,
		})
	}
	if ouo.mutation.IPAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: operation.FieldIPAddress,
		})
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: operation.FieldStatus,
		})
	}
	if value, ok := ouo.mutation.DoneAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operation.FieldDoneAt,
		})
	}
	if ouo.mutation.DoneAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: operation.FieldDoneAt,
		})
	}
	_node = &Operation{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{operation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
