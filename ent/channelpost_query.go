// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"moments/ent/channelpost"
	"moments/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChannelPostQuery is the builder for querying ChannelPost entities.
type ChannelPostQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ChannelPost
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChannelPostQuery builder.
func (cpq *ChannelPostQuery) Where(ps ...predicate.ChannelPost) *ChannelPostQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *ChannelPostQuery) Limit(limit int) *ChannelPostQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *ChannelPostQuery) Offset(offset int) *ChannelPostQuery {
	cpq.offset = &offset
	return cpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpq *ChannelPostQuery) Unique(unique bool) *ChannelPostQuery {
	cpq.unique = &unique
	return cpq
}

// Order adds an order step to the query.
func (cpq *ChannelPostQuery) Order(o ...OrderFunc) *ChannelPostQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// First returns the first ChannelPost entity from the query.
// Returns a *NotFoundError when no ChannelPost was found.
func (cpq *ChannelPostQuery) First(ctx context.Context) (*ChannelPost, error) {
	nodes, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{channelpost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *ChannelPostQuery) FirstX(ctx context.Context) *ChannelPost {
	node, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChannelPost ID from the query.
// Returns a *NotFoundError when no ChannelPost ID was found.
func (cpq *ChannelPostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{channelpost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpq *ChannelPostQuery) FirstIDX(ctx context.Context) int {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChannelPost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChannelPost entity is found.
// Returns a *NotFoundError when no ChannelPost entities are found.
func (cpq *ChannelPostQuery) Only(ctx context.Context) (*ChannelPost, error) {
	nodes, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{channelpost.Label}
	default:
		return nil, &NotSingularError{channelpost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *ChannelPostQuery) OnlyX(ctx context.Context) *ChannelPost {
	node, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChannelPost ID in the query.
// Returns a *NotSingularError when more than one ChannelPost ID is found.
// Returns a *NotFoundError when no entities are found.
func (cpq *ChannelPostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{channelpost.Label}
	default:
		err = &NotSingularError{channelpost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpq *ChannelPostQuery) OnlyIDX(ctx context.Context) int {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChannelPosts.
func (cpq *ChannelPostQuery) All(ctx context.Context) ([]*ChannelPost, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *ChannelPostQuery) AllX(ctx context.Context) []*ChannelPost {
	nodes, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChannelPost IDs.
func (cpq *ChannelPostQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cpq.Select(channelpost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *ChannelPostQuery) IDsX(ctx context.Context) []int {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *ChannelPostQuery) Count(ctx context.Context) (int, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *ChannelPostQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *ChannelPostQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *ChannelPostQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChannelPostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *ChannelPostQuery) Clone() *ChannelPostQuery {
	if cpq == nil {
		return nil
	}
	return &ChannelPostQuery{
		config:     cpq.config,
		limit:      cpq.limit,
		offset:     cpq.offset,
		order:      append([]OrderFunc{}, cpq.order...),
		predicates: append([]predicate.ChannelPost{}, cpq.predicates...),
		// clone intermediate query.
		sql:    cpq.sql.Clone(),
		path:   cpq.path,
		unique: cpq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (cpq *ChannelPostQuery) GroupBy(field string, fields ...string) *ChannelPostGroupBy {
	grbuild := &ChannelPostGroupBy{config: cpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpq.sqlQuery(ctx), nil
	}
	grbuild.label = channelpost.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (cpq *ChannelPostQuery) Select(fields ...string) *ChannelPostSelect {
	cpq.fields = append(cpq.fields, fields...)
	selbuild := &ChannelPostSelect{ChannelPostQuery: cpq}
	selbuild.label = channelpost.Label
	selbuild.flds, selbuild.scan = &cpq.fields, selbuild.Scan
	return selbuild
}

func (cpq *ChannelPostQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpq.fields {
		if !channelpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cpq.path != nil {
		prev, err := cpq.path(ctx)
		if err != nil {
			return err
		}
		cpq.sql = prev
	}
	return nil
}

func (cpq *ChannelPostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChannelPost, error) {
	var (
		nodes = []*ChannelPost{}
		_spec = cpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ChannelPost).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ChannelPost{config: cpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cpq *ChannelPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpq.querySpec()
	_spec.Node.Columns = cpq.fields
	if len(cpq.fields) > 0 {
		_spec.Unique = cpq.unique != nil && *cpq.unique
	}
	return sqlgraph.CountNodes(ctx, cpq.driver, _spec)
}

func (cpq *ChannelPostQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpq *ChannelPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channelpost.Table,
			Columns: channelpost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: channelpost.FieldID,
			},
		},
		From:   cpq.sql,
		Unique: true,
	}
	if unique := cpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channelpost.FieldID)
		for i := range fields {
			if fields[i] != channelpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpq *ChannelPostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(channelpost.Table)
	columns := cpq.fields
	if len(columns) == 0 {
		columns = channelpost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cpq.unique != nil && *cpq.unique {
		selector.Distinct()
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChannelPostGroupBy is the group-by builder for ChannelPost entities.
type ChannelPostGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *ChannelPostGroupBy) Aggregate(fns ...AggregateFunc) *ChannelPostGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpgb *ChannelPostGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpgb.path(ctx)
	if err != nil {
		return err
	}
	cpgb.sql = query
	return cpgb.sqlScan(ctx, v)
}

func (cpgb *ChannelPostGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpgb.fields {
		if !channelpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *ChannelPostGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql.Select()
	aggregation := make([]string, 0, len(cpgb.fns))
	for _, fn := range cpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
		for _, f := range cpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpgb.fields...)...)
}

// ChannelPostSelect is the builder for selecting fields of ChannelPost entities.
type ChannelPostSelect struct {
	*ChannelPostQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cps *ChannelPostSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cps.prepareQuery(ctx); err != nil {
		return err
	}
	cps.sql = cps.ChannelPostQuery.sqlQuery(ctx)
	return cps.sqlScan(ctx, v)
}

func (cps *ChannelPostSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sql.Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
