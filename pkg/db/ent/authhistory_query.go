// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/authhistory"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AuthHistoryQuery is the builder for querying AuthHistory entities.
type AuthHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthHistory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthHistoryQuery builder.
func (ahq *AuthHistoryQuery) Where(ps ...predicate.AuthHistory) *AuthHistoryQuery {
	ahq.predicates = append(ahq.predicates, ps...)
	return ahq
}

// Limit adds a limit step to the query.
func (ahq *AuthHistoryQuery) Limit(limit int) *AuthHistoryQuery {
	ahq.limit = &limit
	return ahq
}

// Offset adds an offset step to the query.
func (ahq *AuthHistoryQuery) Offset(offset int) *AuthHistoryQuery {
	ahq.offset = &offset
	return ahq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ahq *AuthHistoryQuery) Unique(unique bool) *AuthHistoryQuery {
	ahq.unique = &unique
	return ahq
}

// Order adds an order step to the query.
func (ahq *AuthHistoryQuery) Order(o ...OrderFunc) *AuthHistoryQuery {
	ahq.order = append(ahq.order, o...)
	return ahq
}

// First returns the first AuthHistory entity from the query.
// Returns a *NotFoundError when no AuthHistory was found.
func (ahq *AuthHistoryQuery) First(ctx context.Context) (*AuthHistory, error) {
	nodes, err := ahq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ahq *AuthHistoryQuery) FirstX(ctx context.Context) *AuthHistory {
	node, err := ahq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthHistory ID from the query.
// Returns a *NotFoundError when no AuthHistory ID was found.
func (ahq *AuthHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ahq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ahq *AuthHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ahq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthHistory entity is not found.
// Returns a *NotFoundError when no AuthHistory entities are found.
func (ahq *AuthHistoryQuery) Only(ctx context.Context) (*AuthHistory, error) {
	nodes, err := ahq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authhistory.Label}
	default:
		return nil, &NotSingularError{authhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ahq *AuthHistoryQuery) OnlyX(ctx context.Context) *AuthHistory {
	node, err := ahq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthHistory ID in the query.
// Returns a *NotSingularError when exactly one AuthHistory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ahq *AuthHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ahq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = &NotSingularError{authhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ahq *AuthHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ahq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthHistories.
func (ahq *AuthHistoryQuery) All(ctx context.Context) ([]*AuthHistory, error) {
	if err := ahq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ahq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ahq *AuthHistoryQuery) AllX(ctx context.Context) []*AuthHistory {
	nodes, err := ahq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthHistory IDs.
func (ahq *AuthHistoryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ahq.Select(authhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ahq *AuthHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ahq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ahq *AuthHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := ahq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ahq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ahq *AuthHistoryQuery) CountX(ctx context.Context) int {
	count, err := ahq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ahq *AuthHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := ahq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ahq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ahq *AuthHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ahq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ahq *AuthHistoryQuery) Clone() *AuthHistoryQuery {
	if ahq == nil {
		return nil
	}
	return &AuthHistoryQuery{
		config:     ahq.config,
		limit:      ahq.limit,
		offset:     ahq.offset,
		order:      append([]OrderFunc{}, ahq.order...),
		predicates: append([]predicate.AuthHistory{}, ahq.predicates...),
		// clone intermediate query.
		sql:  ahq.sql.Clone(),
		path: ahq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthHistory.Query().
//		GroupBy(authhistory.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ahq *AuthHistoryQuery) GroupBy(field string, fields ...string) *AuthHistoryGroupBy {
	group := &AuthHistoryGroupBy{config: ahq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ahq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ahq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AuthHistory.Query().
//		Select(authhistory.FieldAppID).
//		Scan(ctx, &v)
//
func (ahq *AuthHistoryQuery) Select(fields ...string) *AuthHistorySelect {
	ahq.fields = append(ahq.fields, fields...)
	return &AuthHistorySelect{AuthHistoryQuery: ahq}
}

func (ahq *AuthHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ahq.fields {
		if !authhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ahq.path != nil {
		prev, err := ahq.path(ctx)
		if err != nil {
			return err
		}
		ahq.sql = prev
	}
	return nil
}

func (ahq *AuthHistoryQuery) sqlAll(ctx context.Context) ([]*AuthHistory, error) {
	var (
		nodes = []*AuthHistory{}
		_spec = ahq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthHistory{config: ahq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ahq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ahq *AuthHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ahq.querySpec()
	_spec.Node.Columns = ahq.fields
	if len(ahq.fields) > 0 {
		_spec.Unique = ahq.unique != nil && *ahq.unique
	}
	return sqlgraph.CountNodes(ctx, ahq.driver, _spec)
}

func (ahq *AuthHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ahq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ahq *AuthHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authhistory.Table,
			Columns: authhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authhistory.FieldID,
			},
		},
		From:   ahq.sql,
		Unique: true,
	}
	if unique := ahq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ahq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authhistory.FieldID)
		for i := range fields {
			if fields[i] != authhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ahq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ahq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ahq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ahq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ahq *AuthHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ahq.driver.Dialect())
	t1 := builder.Table(authhistory.Table)
	columns := ahq.fields
	if len(columns) == 0 {
		columns = authhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ahq.sql != nil {
		selector = ahq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ahq.unique != nil && *ahq.unique {
		selector.Distinct()
	}
	for _, p := range ahq.predicates {
		p(selector)
	}
	for _, p := range ahq.order {
		p(selector)
	}
	if offset := ahq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ahq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthHistoryGroupBy is the group-by builder for AuthHistory entities.
type AuthHistoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ahgb *AuthHistoryGroupBy) Aggregate(fns ...AggregateFunc) *AuthHistoryGroupBy {
	ahgb.fns = append(ahgb.fns, fns...)
	return ahgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ahgb *AuthHistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ahgb.path(ctx)
	if err != nil {
		return err
	}
	ahgb.sql = query
	return ahgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ahgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ahgb.fields) > 1 {
		return nil, errors.New("ent: AuthHistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := ahgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ahgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) StringX(ctx context.Context) string {
	v, err := ahgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ahgb.fields) > 1 {
		return nil, errors.New("ent: AuthHistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := ahgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ahgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) IntX(ctx context.Context) int {
	v, err := ahgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ahgb.fields) > 1 {
		return nil, errors.New("ent: AuthHistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ahgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ahgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ahgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ahgb.fields) > 1 {
		return nil, errors.New("ent: AuthHistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ahgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ahgb *AuthHistoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ahgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ahgb *AuthHistoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := ahgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ahgb *AuthHistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ahgb.fields {
		if !authhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ahgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ahgb *AuthHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := ahgb.sql.Select()
	aggregation := make([]string, 0, len(ahgb.fns))
	for _, fn := range ahgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ahgb.fields)+len(ahgb.fns))
		for _, f := range ahgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ahgb.fields...)...)
}

// AuthHistorySelect is the builder for selecting fields of AuthHistory entities.
type AuthHistorySelect struct {
	*AuthHistoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ahs *AuthHistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ahs.prepareQuery(ctx); err != nil {
		return err
	}
	ahs.sql = ahs.AuthHistoryQuery.sqlQuery(ctx)
	return ahs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ahs *AuthHistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ahs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ahs.fields) > 1 {
		return nil, errors.New("ent: AuthHistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ahs *AuthHistorySelect) StringsX(ctx context.Context) []string {
	v, err := ahs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ahs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ahs *AuthHistorySelect) StringX(ctx context.Context) string {
	v, err := ahs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ahs.fields) > 1 {
		return nil, errors.New("ent: AuthHistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ahs *AuthHistorySelect) IntsX(ctx context.Context) []int {
	v, err := ahs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ahs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ahs *AuthHistorySelect) IntX(ctx context.Context) int {
	v, err := ahs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ahs.fields) > 1 {
		return nil, errors.New("ent: AuthHistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ahs *AuthHistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ahs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ahs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ahs *AuthHistorySelect) Float64X(ctx context.Context) float64 {
	v, err := ahs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ahs.fields) > 1 {
		return nil, errors.New("ent: AuthHistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ahs *AuthHistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := ahs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ahs *AuthHistorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ahs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authhistory.Label}
	default:
		err = fmt.Errorf("ent: AuthHistorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ahs *AuthHistorySelect) BoolX(ctx context.Context) bool {
	v, err := ahs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ahs *AuthHistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ahs.sql.Query()
	if err := ahs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
