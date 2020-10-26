// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblquery"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblresponse"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/ip"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// DNSBLQueryUpdate is the builder for updating DNSBLQuery entities.
type DNSBLQueryUpdate struct {
	config
	hooks    []Hook
	mutation *DNSBLQueryMutation
}

// Where adds a new predicate for the builder.
func (dqu *DNSBLQueryUpdate) Where(ps ...predicate.DNSBLQuery) *DNSBLQueryUpdate {
	dqu.mutation.predicates = append(dqu.mutation.predicates, ps...)
	return dqu
}

// SetCreatedAt sets the created_at field.
func (dqu *DNSBLQueryUpdate) SetCreatedAt(t time.Time) *DNSBLQueryUpdate {
	dqu.mutation.SetCreatedAt(t)
	return dqu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (dqu *DNSBLQueryUpdate) SetNillableCreatedAt(t *time.Time) *DNSBLQueryUpdate {
	if t != nil {
		dqu.SetCreatedAt(*t)
	}
	return dqu
}

// SetUpdatedAt sets the updated_at field.
func (dqu *DNSBLQueryUpdate) SetUpdatedAt(t time.Time) *DNSBLQueryUpdate {
	dqu.mutation.SetUpdatedAt(t)
	return dqu
}

// AddResponseIDs adds the responses edge to DNSBLResponse by ids.
func (dqu *DNSBLQueryUpdate) AddResponseIDs(ids ...uuid.UUID) *DNSBLQueryUpdate {
	dqu.mutation.AddResponseIDs(ids...)
	return dqu
}

// AddResponses adds the responses edges to DNSBLResponse.
func (dqu *DNSBLQueryUpdate) AddResponses(d ...*DNSBLResponse) *DNSBLQueryUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dqu.AddResponseIDs(ids...)
}

// SetIPAddressID sets the ip_address edge to IP by id.
func (dqu *DNSBLQueryUpdate) SetIPAddressID(id uuid.UUID) *DNSBLQueryUpdate {
	dqu.mutation.SetIPAddressID(id)
	return dqu
}

// SetIPAddress sets the ip_address edge to IP.
func (dqu *DNSBLQueryUpdate) SetIPAddress(i *IP) *DNSBLQueryUpdate {
	return dqu.SetIPAddressID(i.ID)
}

// Mutation returns the DNSBLQueryMutation object of the builder.
func (dqu *DNSBLQueryUpdate) Mutation() *DNSBLQueryMutation {
	return dqu.mutation
}

// ClearResponses clears all "responses" edges to type DNSBLResponse.
func (dqu *DNSBLQueryUpdate) ClearResponses() *DNSBLQueryUpdate {
	dqu.mutation.ClearResponses()
	return dqu
}

// RemoveResponseIDs removes the responses edge to DNSBLResponse by ids.
func (dqu *DNSBLQueryUpdate) RemoveResponseIDs(ids ...uuid.UUID) *DNSBLQueryUpdate {
	dqu.mutation.RemoveResponseIDs(ids...)
	return dqu
}

// RemoveResponses removes responses edges to DNSBLResponse.
func (dqu *DNSBLQueryUpdate) RemoveResponses(d ...*DNSBLResponse) *DNSBLQueryUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dqu.RemoveResponseIDs(ids...)
}

// ClearIPAddress clears the "ip_address" edge to type IP.
func (dqu *DNSBLQueryUpdate) ClearIPAddress() *DNSBLQueryUpdate {
	dqu.mutation.ClearIPAddress()
	return dqu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (dqu *DNSBLQueryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	dqu.defaults()
	if len(dqu.hooks) == 0 {
		if err = dqu.check(); err != nil {
			return 0, err
		}
		affected, err = dqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSBLQueryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dqu.check(); err != nil {
				return 0, err
			}
			dqu.mutation = mutation
			affected, err = dqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dqu.hooks) - 1; i >= 0; i-- {
			mut = dqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dqu *DNSBLQueryUpdate) SaveX(ctx context.Context) int {
	affected, err := dqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dqu *DNSBLQueryUpdate) Exec(ctx context.Context) error {
	_, err := dqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dqu *DNSBLQueryUpdate) ExecX(ctx context.Context) {
	if err := dqu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dqu *DNSBLQueryUpdate) defaults() {
	if _, ok := dqu.mutation.UpdatedAt(); !ok {
		v := dnsblquery.UpdateDefaultUpdatedAt()
		dqu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dqu *DNSBLQueryUpdate) check() error {
	if _, ok := dqu.mutation.IPAddressID(); dqu.mutation.IPAddressCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"ip_address\"")
	}
	return nil
}

func (dqu *DNSBLQueryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsblquery.Table,
			Columns: dnsblquery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dnsblquery.FieldID,
			},
		},
	}
	if ps := dqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dqu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsblquery.FieldCreatedAt,
		})
	}
	if value, ok := dqu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsblquery.FieldUpdatedAt,
		})
	}
	if dqu.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsblquery.ResponsesTable,
			Columns: []string{dnsblquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblresponse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dqu.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !dqu.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsblquery.ResponsesTable,
			Columns: []string{dnsblquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dqu.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsblquery.ResponsesTable,
			Columns: []string{dnsblquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dqu.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dnsblquery.IPAddressTable,
			Columns: []string{dnsblquery.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dqu.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dnsblquery.IPAddressTable,
			Columns: []string{dnsblquery.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsblquery.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DNSBLQueryUpdateOne is the builder for updating a single DNSBLQuery entity.
type DNSBLQueryUpdateOne struct {
	config
	hooks    []Hook
	mutation *DNSBLQueryMutation
}

// SetCreatedAt sets the created_at field.
func (dquo *DNSBLQueryUpdateOne) SetCreatedAt(t time.Time) *DNSBLQueryUpdateOne {
	dquo.mutation.SetCreatedAt(t)
	return dquo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (dquo *DNSBLQueryUpdateOne) SetNillableCreatedAt(t *time.Time) *DNSBLQueryUpdateOne {
	if t != nil {
		dquo.SetCreatedAt(*t)
	}
	return dquo
}

// SetUpdatedAt sets the updated_at field.
func (dquo *DNSBLQueryUpdateOne) SetUpdatedAt(t time.Time) *DNSBLQueryUpdateOne {
	dquo.mutation.SetUpdatedAt(t)
	return dquo
}

// AddResponseIDs adds the responses edge to DNSBLResponse by ids.
func (dquo *DNSBLQueryUpdateOne) AddResponseIDs(ids ...uuid.UUID) *DNSBLQueryUpdateOne {
	dquo.mutation.AddResponseIDs(ids...)
	return dquo
}

// AddResponses adds the responses edges to DNSBLResponse.
func (dquo *DNSBLQueryUpdateOne) AddResponses(d ...*DNSBLResponse) *DNSBLQueryUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dquo.AddResponseIDs(ids...)
}

// SetIPAddressID sets the ip_address edge to IP by id.
func (dquo *DNSBLQueryUpdateOne) SetIPAddressID(id uuid.UUID) *DNSBLQueryUpdateOne {
	dquo.mutation.SetIPAddressID(id)
	return dquo
}

// SetIPAddress sets the ip_address edge to IP.
func (dquo *DNSBLQueryUpdateOne) SetIPAddress(i *IP) *DNSBLQueryUpdateOne {
	return dquo.SetIPAddressID(i.ID)
}

// Mutation returns the DNSBLQueryMutation object of the builder.
func (dquo *DNSBLQueryUpdateOne) Mutation() *DNSBLQueryMutation {
	return dquo.mutation
}

// ClearResponses clears all "responses" edges to type DNSBLResponse.
func (dquo *DNSBLQueryUpdateOne) ClearResponses() *DNSBLQueryUpdateOne {
	dquo.mutation.ClearResponses()
	return dquo
}

// RemoveResponseIDs removes the responses edge to DNSBLResponse by ids.
func (dquo *DNSBLQueryUpdateOne) RemoveResponseIDs(ids ...uuid.UUID) *DNSBLQueryUpdateOne {
	dquo.mutation.RemoveResponseIDs(ids...)
	return dquo
}

// RemoveResponses removes responses edges to DNSBLResponse.
func (dquo *DNSBLQueryUpdateOne) RemoveResponses(d ...*DNSBLResponse) *DNSBLQueryUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dquo.RemoveResponseIDs(ids...)
}

// ClearIPAddress clears the "ip_address" edge to type IP.
func (dquo *DNSBLQueryUpdateOne) ClearIPAddress() *DNSBLQueryUpdateOne {
	dquo.mutation.ClearIPAddress()
	return dquo
}

// Save executes the query and returns the updated entity.
func (dquo *DNSBLQueryUpdateOne) Save(ctx context.Context) (*DNSBLQuery, error) {
	var (
		err  error
		node *DNSBLQuery
	)
	dquo.defaults()
	if len(dquo.hooks) == 0 {
		if err = dquo.check(); err != nil {
			return nil, err
		}
		node, err = dquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSBLQueryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dquo.check(); err != nil {
				return nil, err
			}
			dquo.mutation = mutation
			node, err = dquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dquo.hooks) - 1; i >= 0; i-- {
			mut = dquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dquo *DNSBLQueryUpdateOne) SaveX(ctx context.Context) *DNSBLQuery {
	node, err := dquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dquo *DNSBLQueryUpdateOne) Exec(ctx context.Context) error {
	_, err := dquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dquo *DNSBLQueryUpdateOne) ExecX(ctx context.Context) {
	if err := dquo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dquo *DNSBLQueryUpdateOne) defaults() {
	if _, ok := dquo.mutation.UpdatedAt(); !ok {
		v := dnsblquery.UpdateDefaultUpdatedAt()
		dquo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dquo *DNSBLQueryUpdateOne) check() error {
	if _, ok := dquo.mutation.IPAddressID(); dquo.mutation.IPAddressCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"ip_address\"")
	}
	return nil
}

func (dquo *DNSBLQueryUpdateOne) sqlSave(ctx context.Context) (_node *DNSBLQuery, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsblquery.Table,
			Columns: dnsblquery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dnsblquery.FieldID,
			},
		},
	}
	id, ok := dquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DNSBLQuery.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := dquo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsblquery.FieldCreatedAt,
		})
	}
	if value, ok := dquo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsblquery.FieldUpdatedAt,
		})
	}
	if dquo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsblquery.ResponsesTable,
			Columns: []string{dnsblquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblresponse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dquo.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !dquo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsblquery.ResponsesTable,
			Columns: []string{dnsblquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dquo.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsblquery.ResponsesTable,
			Columns: []string{dnsblquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dquo.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dnsblquery.IPAddressTable,
			Columns: []string{dnsblquery.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dquo.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dnsblquery.IPAddressTable,
			Columns: []string{dnsblquery.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DNSBLQuery{config: dquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, dquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsblquery.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
