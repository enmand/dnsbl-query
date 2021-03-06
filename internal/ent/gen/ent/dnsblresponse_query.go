// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblquery"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblresponse"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// DNSBLResponseQuery is the builder for querying DNSBLResponse entities.
type DNSBLResponseQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.DNSBLResponse
	// eager-loading edges.
	withQuery *DNSBLQueryQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (drq *DNSBLResponseQuery) Where(ps ...predicate.DNSBLResponse) *DNSBLResponseQuery {
	drq.predicates = append(drq.predicates, ps...)
	return drq
}

// Limit adds a limit step to the query.
func (drq *DNSBLResponseQuery) Limit(limit int) *DNSBLResponseQuery {
	drq.limit = &limit
	return drq
}

// Offset adds an offset step to the query.
func (drq *DNSBLResponseQuery) Offset(offset int) *DNSBLResponseQuery {
	drq.offset = &offset
	return drq
}

// Order adds an order step to the query.
func (drq *DNSBLResponseQuery) Order(o ...OrderFunc) *DNSBLResponseQuery {
	drq.order = append(drq.order, o...)
	return drq
}

// QueryQuery chains the current query on the query edge.
func (drq *DNSBLResponseQuery) QueryQuery() *DNSBLQueryQuery {
	query := &DNSBLQueryQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsblresponse.Table, dnsblresponse.FieldID, selector),
			sqlgraph.To(dnsblquery.Table, dnsblquery.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dnsblresponse.QueryTable, dnsblresponse.QueryColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DNSBLResponse entity in the query. Returns *NotFoundError when no dnsblresponse was found.
func (drq *DNSBLResponseQuery) First(ctx context.Context) (*DNSBLResponse, error) {
	nodes, err := drq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dnsblresponse.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drq *DNSBLResponseQuery) FirstX(ctx context.Context) *DNSBLResponse {
	node, err := drq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DNSBLResponse id in the query. Returns *NotFoundError when no id was found.
func (drq *DNSBLResponseQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = drq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dnsblresponse.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (drq *DNSBLResponseQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := drq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only DNSBLResponse entity in the query, returns an error if not exactly one entity was returned.
func (drq *DNSBLResponseQuery) Only(ctx context.Context) (*DNSBLResponse, error) {
	nodes, err := drq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dnsblresponse.Label}
	default:
		return nil, &NotSingularError{dnsblresponse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drq *DNSBLResponseQuery) OnlyX(ctx context.Context) *DNSBLResponse {
	node, err := drq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only DNSBLResponse id in the query, returns an error if not exactly one id was returned.
func (drq *DNSBLResponseQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = drq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = &NotSingularError{dnsblresponse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drq *DNSBLResponseQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := drq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DNSBLResponses.
func (drq *DNSBLResponseQuery) All(ctx context.Context) ([]*DNSBLResponse, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return drq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (drq *DNSBLResponseQuery) AllX(ctx context.Context) []*DNSBLResponse {
	nodes, err := drq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DNSBLResponse ids.
func (drq *DNSBLResponseQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := drq.Select(dnsblresponse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drq *DNSBLResponseQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := drq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drq *DNSBLResponseQuery) Count(ctx context.Context) (int, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return drq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (drq *DNSBLResponseQuery) CountX(ctx context.Context) int {
	count, err := drq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drq *DNSBLResponseQuery) Exist(ctx context.Context) (bool, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return drq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (drq *DNSBLResponseQuery) ExistX(ctx context.Context) bool {
	exist, err := drq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drq *DNSBLResponseQuery) Clone() *DNSBLResponseQuery {
	return &DNSBLResponseQuery{
		config:     drq.config,
		limit:      drq.limit,
		offset:     drq.offset,
		order:      append([]OrderFunc{}, drq.order...),
		unique:     append([]string{}, drq.unique...),
		predicates: append([]predicate.DNSBLResponse{}, drq.predicates...),
		// clone intermediate query.
		sql:  drq.sql.Clone(),
		path: drq.path,
	}
}

//  WithQuery tells the query-builder to eager-loads the nodes that are connected to
// the "query" edge. The optional arguments used to configure the query builder of the edge.
func (drq *DNSBLResponseQuery) WithQuery(opts ...func(*DNSBLQueryQuery)) *DNSBLResponseQuery {
	query := &DNSBLQueryQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withQuery = query
	return drq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DNSBLResponse.Query().
//		GroupBy(dnsblresponse.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (drq *DNSBLResponseQuery) GroupBy(field string, fields ...string) *DNSBLResponseGroupBy {
	group := &DNSBLResponseGroupBy{config: drq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DNSBLResponse.Query().
//		Select(dnsblresponse.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (drq *DNSBLResponseQuery) Select(field string, fields ...string) *DNSBLResponseSelect {
	selector := &DNSBLResponseSelect{config: drq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(), nil
	}
	return selector
}

func (drq *DNSBLResponseQuery) prepareQuery(ctx context.Context) error {
	if drq.path != nil {
		prev, err := drq.path(ctx)
		if err != nil {
			return err
		}
		drq.sql = prev
	}
	return nil
}

func (drq *DNSBLResponseQuery) sqlAll(ctx context.Context) ([]*DNSBLResponse, error) {
	var (
		nodes       = []*DNSBLResponse{}
		withFKs     = drq.withFKs
		_spec       = drq.querySpec()
		loadedTypes = [1]bool{
			drq.withQuery != nil,
		}
	)
	if drq.withQuery != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dnsblresponse.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &DNSBLResponse{config: drq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, drq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := drq.withQuery; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*DNSBLResponse)
		for i := range nodes {
			if fk := nodes[i].dnsbl_query_responses; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(dnsblquery.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dnsbl_query_responses" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Query = n
			}
		}
	}

	return nodes, nil
}

func (drq *DNSBLResponseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drq.querySpec()
	return sqlgraph.CountNodes(ctx, drq.driver, _spec)
}

func (drq *DNSBLResponseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := drq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (drq *DNSBLResponseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsblresponse.Table,
			Columns: dnsblresponse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dnsblresponse.FieldID,
			},
		},
		From:   drq.sql,
		Unique: true,
	}
	if ps := drq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, dnsblresponse.ValidColumn)
			}
		}
	}
	return _spec
}

func (drq *DNSBLResponseQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(drq.driver.Dialect())
	t1 := builder.Table(dnsblresponse.Table)
	selector := builder.Select(t1.Columns(dnsblresponse.Columns...)...).From(t1)
	if drq.sql != nil {
		selector = drq.sql
		selector.Select(selector.Columns(dnsblresponse.Columns...)...)
	}
	for _, p := range drq.predicates {
		p(selector)
	}
	for _, p := range drq.order {
		p(selector, dnsblresponse.ValidColumn)
	}
	if offset := drq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DNSBLResponseGroupBy is the builder for group-by DNSBLResponse entities.
type DNSBLResponseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drgb *DNSBLResponseGroupBy) Aggregate(fns ...AggregateFunc) *DNSBLResponseGroupBy {
	drgb.fns = append(drgb.fns, fns...)
	return drgb
}

// Scan applies the group-by query and scan the result into the given value.
func (drgb *DNSBLResponseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := drgb.path(ctx)
	if err != nil {
		return err
	}
	drgb.sql = query
	return drgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := drgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) StringsX(ctx context.Context) []string {
	v, err := drgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) StringX(ctx context.Context) string {
	v, err := drgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) IntsX(ctx context.Context) []int {
	v, err := drgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) IntX(ctx context.Context) int {
	v, err := drgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := drgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := drgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := drgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (drgb *DNSBLResponseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drgb *DNSBLResponseGroupBy) BoolX(ctx context.Context) bool {
	v, err := drgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drgb *DNSBLResponseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range drgb.fields {
		if !dnsblresponse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := drgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drgb *DNSBLResponseGroupBy) sqlQuery() *sql.Selector {
	selector := drgb.sql
	columns := make([]string, 0, len(drgb.fields)+len(drgb.fns))
	columns = append(columns, drgb.fields...)
	for _, fn := range drgb.fns {
		columns = append(columns, fn(selector, dnsblresponse.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(drgb.fields...)
}

// DNSBLResponseSelect is the builder for select fields of DNSBLResponse entities.
type DNSBLResponseSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (drs *DNSBLResponseSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := drs.path(ctx)
	if err != nil {
		return err
	}
	drs.sql = query
	return drs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drs *DNSBLResponseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := drs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drs *DNSBLResponseSelect) StringsX(ctx context.Context) []string {
	v, err := drs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drs *DNSBLResponseSelect) StringX(ctx context.Context) string {
	v, err := drs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drs *DNSBLResponseSelect) IntsX(ctx context.Context) []int {
	v, err := drs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drs *DNSBLResponseSelect) IntX(ctx context.Context) int {
	v, err := drs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drs *DNSBLResponseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := drs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drs *DNSBLResponseSelect) Float64X(ctx context.Context) float64 {
	v, err := drs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSBLResponseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drs *DNSBLResponseSelect) BoolsX(ctx context.Context) []bool {
	v, err := drs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (drs *DNSBLResponseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsblresponse.Label}
	default:
		err = fmt.Errorf("ent: DNSBLResponseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drs *DNSBLResponseSelect) BoolX(ctx context.Context) bool {
	v, err := drs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drs *DNSBLResponseSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range drs.fields {
		if !dnsblresponse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := drs.sqlQuery().Query()
	if err := drs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drs *DNSBLResponseSelect) sqlQuery() sql.Querier {
	selector := drs.sql
	selector.Select(selector.Columns(drs.fields...)...)
	return selector
}
