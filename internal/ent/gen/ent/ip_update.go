// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblquery"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/ip"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// IPUpdate is the builder for updating IP entities.
type IPUpdate struct {
	config
	hooks    []Hook
	mutation *IPMutation
}

// Where adds a new predicate for the builder.
func (iu *IPUpdate) Where(ps ...predicate.IP) *IPUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetCreatedAt sets the created_at field.
func (iu *IPUpdate) SetCreatedAt(t time.Time) *IPUpdate {
	iu.mutation.SetCreatedAt(t)
	return iu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (iu *IPUpdate) SetNillableCreatedAt(t *time.Time) *IPUpdate {
	if t != nil {
		iu.SetCreatedAt(*t)
	}
	return iu
}

// SetUpdatedAt sets the updated_at field.
func (iu *IPUpdate) SetUpdatedAt(t time.Time) *IPUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetIPAddress sets the ip_address field.
func (iu *IPUpdate) SetIPAddress(s string) *IPUpdate {
	iu.mutation.SetIPAddress(s)
	return iu
}

// AddQueryIDs adds the queries edge to DNSBLQuery by ids.
func (iu *IPUpdate) AddQueryIDs(ids ...uuid.UUID) *IPUpdate {
	iu.mutation.AddQueryIDs(ids...)
	return iu
}

// AddQueries adds the queries edges to DNSBLQuery.
func (iu *IPUpdate) AddQueries(d ...*DNSBLQuery) *IPUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.AddQueryIDs(ids...)
}

// Mutation returns the IPMutation object of the builder.
func (iu *IPUpdate) Mutation() *IPMutation {
	return iu.mutation
}

// ClearQueries clears all "queries" edges to type DNSBLQuery.
func (iu *IPUpdate) ClearQueries() *IPUpdate {
	iu.mutation.ClearQueries()
	return iu
}

// RemoveQueryIDs removes the queries edge to DNSBLQuery by ids.
func (iu *IPUpdate) RemoveQueryIDs(ids ...uuid.UUID) *IPUpdate {
	iu.mutation.RemoveQueryIDs(ids...)
	return iu
}

// RemoveQueries removes queries edges to DNSBLQuery.
func (iu *IPUpdate) RemoveQueries(d ...*DNSBLQuery) *IPUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.RemoveQueryIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *IPUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IPUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IPUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IPUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IPUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := ip.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

func (iu *IPUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ip.Table,
			Columns: ip.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ip.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ip.FieldCreatedAt,
		})
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ip.FieldUpdatedAt,
		})
	}
	if value, ok := iu.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ip.FieldIPAddress,
		})
	}
	if iu.mutation.QueriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ip.QueriesTable,
			Columns: []string{ip.QueriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblquery.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedQueriesIDs(); len(nodes) > 0 && !iu.mutation.QueriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ip.QueriesTable,
			Columns: []string{ip.QueriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblquery.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.QueriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ip.QueriesTable,
			Columns: []string{ip.QueriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblquery.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ip.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IPUpdateOne is the builder for updating a single IP entity.
type IPUpdateOne struct {
	config
	hooks    []Hook
	mutation *IPMutation
}

// SetCreatedAt sets the created_at field.
func (iuo *IPUpdateOne) SetCreatedAt(t time.Time) *IPUpdateOne {
	iuo.mutation.SetCreatedAt(t)
	return iuo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (iuo *IPUpdateOne) SetNillableCreatedAt(t *time.Time) *IPUpdateOne {
	if t != nil {
		iuo.SetCreatedAt(*t)
	}
	return iuo
}

// SetUpdatedAt sets the updated_at field.
func (iuo *IPUpdateOne) SetUpdatedAt(t time.Time) *IPUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetIPAddress sets the ip_address field.
func (iuo *IPUpdateOne) SetIPAddress(s string) *IPUpdateOne {
	iuo.mutation.SetIPAddress(s)
	return iuo
}

// AddQueryIDs adds the queries edge to DNSBLQuery by ids.
func (iuo *IPUpdateOne) AddQueryIDs(ids ...uuid.UUID) *IPUpdateOne {
	iuo.mutation.AddQueryIDs(ids...)
	return iuo
}

// AddQueries adds the queries edges to DNSBLQuery.
func (iuo *IPUpdateOne) AddQueries(d ...*DNSBLQuery) *IPUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.AddQueryIDs(ids...)
}

// Mutation returns the IPMutation object of the builder.
func (iuo *IPUpdateOne) Mutation() *IPMutation {
	return iuo.mutation
}

// ClearQueries clears all "queries" edges to type DNSBLQuery.
func (iuo *IPUpdateOne) ClearQueries() *IPUpdateOne {
	iuo.mutation.ClearQueries()
	return iuo
}

// RemoveQueryIDs removes the queries edge to DNSBLQuery by ids.
func (iuo *IPUpdateOne) RemoveQueryIDs(ids ...uuid.UUID) *IPUpdateOne {
	iuo.mutation.RemoveQueryIDs(ids...)
	return iuo
}

// RemoveQueries removes queries edges to DNSBLQuery.
func (iuo *IPUpdateOne) RemoveQueries(d ...*DNSBLQuery) *IPUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.RemoveQueryIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (iuo *IPUpdateOne) Save(ctx context.Context) (*IP, error) {
	var (
		err  error
		node *IP
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IPUpdateOne) SaveX(ctx context.Context) *IP {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IPUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IPUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IPUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := ip.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

func (iuo *IPUpdateOne) sqlSave(ctx context.Context) (_node *IP, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ip.Table,
			Columns: ip.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ip.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing IP.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ip.FieldCreatedAt,
		})
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ip.FieldUpdatedAt,
		})
	}
	if value, ok := iuo.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ip.FieldIPAddress,
		})
	}
	if iuo.mutation.QueriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ip.QueriesTable,
			Columns: []string{ip.QueriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblquery.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedQueriesIDs(); len(nodes) > 0 && !iuo.mutation.QueriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ip.QueriesTable,
			Columns: []string{ip.QueriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblquery.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.QueriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ip.QueriesTable,
			Columns: []string{ip.QueriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsblquery.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IP{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ip.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
