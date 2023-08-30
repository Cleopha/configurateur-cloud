// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/loadbalancer"
	"backend/ent/panierloadbalancer"
	"backend/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoadBalancerQuery is the builder for querying LoadBalancer entities.
type LoadBalancerQuery struct {
	config
	ctx                    *QueryContext
	order                  []loadbalancer.OrderOption
	inters                 []Interceptor
	predicates             []predicate.LoadBalancer
	withPanierLoadBalancer *PanierLoadBalancerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoadBalancerQuery builder.
func (lbq *LoadBalancerQuery) Where(ps ...predicate.LoadBalancer) *LoadBalancerQuery {
	lbq.predicates = append(lbq.predicates, ps...)
	return lbq
}

// Limit the number of records to be returned by this query.
func (lbq *LoadBalancerQuery) Limit(limit int) *LoadBalancerQuery {
	lbq.ctx.Limit = &limit
	return lbq
}

// Offset to start from.
func (lbq *LoadBalancerQuery) Offset(offset int) *LoadBalancerQuery {
	lbq.ctx.Offset = &offset
	return lbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lbq *LoadBalancerQuery) Unique(unique bool) *LoadBalancerQuery {
	lbq.ctx.Unique = &unique
	return lbq
}

// Order specifies how the records should be ordered.
func (lbq *LoadBalancerQuery) Order(o ...loadbalancer.OrderOption) *LoadBalancerQuery {
	lbq.order = append(lbq.order, o...)
	return lbq
}

// QueryPanierLoadBalancer chains the current query on the "panierLoadBalancer" edge.
func (lbq *LoadBalancerQuery) QueryPanierLoadBalancer() *PanierLoadBalancerQuery {
	query := (&PanierLoadBalancerClient{config: lbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loadbalancer.Table, loadbalancer.FieldID, selector),
			sqlgraph.To(panierloadbalancer.Table, panierloadbalancer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, loadbalancer.PanierLoadBalancerTable, loadbalancer.PanierLoadBalancerColumn),
		)
		fromU = sqlgraph.SetNeighbors(lbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LoadBalancer entity from the query.
// Returns a *NotFoundError when no LoadBalancer was found.
func (lbq *LoadBalancerQuery) First(ctx context.Context) (*LoadBalancer, error) {
	nodes, err := lbq.Limit(1).All(setContextOp(ctx, lbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loadbalancer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lbq *LoadBalancerQuery) FirstX(ctx context.Context) *LoadBalancer {
	node, err := lbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoadBalancer ID from the query.
// Returns a *NotFoundError when no LoadBalancer ID was found.
func (lbq *LoadBalancerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lbq.Limit(1).IDs(setContextOp(ctx, lbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loadbalancer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lbq *LoadBalancerQuery) FirstIDX(ctx context.Context) int {
	id, err := lbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoadBalancer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoadBalancer entity is found.
// Returns a *NotFoundError when no LoadBalancer entities are found.
func (lbq *LoadBalancerQuery) Only(ctx context.Context) (*LoadBalancer, error) {
	nodes, err := lbq.Limit(2).All(setContextOp(ctx, lbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loadbalancer.Label}
	default:
		return nil, &NotSingularError{loadbalancer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lbq *LoadBalancerQuery) OnlyX(ctx context.Context) *LoadBalancer {
	node, err := lbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoadBalancer ID in the query.
// Returns a *NotSingularError when more than one LoadBalancer ID is found.
// Returns a *NotFoundError when no entities are found.
func (lbq *LoadBalancerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lbq.Limit(2).IDs(setContextOp(ctx, lbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loadbalancer.Label}
	default:
		err = &NotSingularError{loadbalancer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lbq *LoadBalancerQuery) OnlyIDX(ctx context.Context) int {
	id, err := lbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoadBalancers.
func (lbq *LoadBalancerQuery) All(ctx context.Context) ([]*LoadBalancer, error) {
	ctx = setContextOp(ctx, lbq.ctx, "All")
	if err := lbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoadBalancer, *LoadBalancerQuery]()
	return withInterceptors[[]*LoadBalancer](ctx, lbq, qr, lbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lbq *LoadBalancerQuery) AllX(ctx context.Context) []*LoadBalancer {
	nodes, err := lbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoadBalancer IDs.
func (lbq *LoadBalancerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lbq.ctx.Unique == nil && lbq.path != nil {
		lbq.Unique(true)
	}
	ctx = setContextOp(ctx, lbq.ctx, "IDs")
	if err = lbq.Select(loadbalancer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lbq *LoadBalancerQuery) IDsX(ctx context.Context) []int {
	ids, err := lbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lbq *LoadBalancerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lbq.ctx, "Count")
	if err := lbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lbq, querierCount[*LoadBalancerQuery](), lbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lbq *LoadBalancerQuery) CountX(ctx context.Context) int {
	count, err := lbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lbq *LoadBalancerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lbq.ctx, "Exist")
	switch _, err := lbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lbq *LoadBalancerQuery) ExistX(ctx context.Context) bool {
	exist, err := lbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoadBalancerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lbq *LoadBalancerQuery) Clone() *LoadBalancerQuery {
	if lbq == nil {
		return nil
	}
	return &LoadBalancerQuery{
		config:                 lbq.config,
		ctx:                    lbq.ctx.Clone(),
		order:                  append([]loadbalancer.OrderOption{}, lbq.order...),
		inters:                 append([]Interceptor{}, lbq.inters...),
		predicates:             append([]predicate.LoadBalancer{}, lbq.predicates...),
		withPanierLoadBalancer: lbq.withPanierLoadBalancer.Clone(),
		// clone intermediate query.
		sql:  lbq.sql.Clone(),
		path: lbq.path,
	}
}

// WithPanierLoadBalancer tells the query-builder to eager-load the nodes that are connected to
// the "panierLoadBalancer" edge. The optional arguments are used to configure the query builder of the edge.
func (lbq *LoadBalancerQuery) WithPanierLoadBalancer(opts ...func(*PanierLoadBalancerQuery)) *LoadBalancerQuery {
	query := (&PanierLoadBalancerClient{config: lbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lbq.withPanierLoadBalancer = query
	return lbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LoadBalancer.Query().
//		GroupBy(loadbalancer.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lbq *LoadBalancerQuery) GroupBy(field string, fields ...string) *LoadBalancerGroupBy {
	lbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoadBalancerGroupBy{build: lbq}
	grbuild.flds = &lbq.ctx.Fields
	grbuild.label = loadbalancer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//	}
//
//	client.LoadBalancer.Query().
//		Select(loadbalancer.FieldName).
//		Scan(ctx, &v)
func (lbq *LoadBalancerQuery) Select(fields ...string) *LoadBalancerSelect {
	lbq.ctx.Fields = append(lbq.ctx.Fields, fields...)
	sbuild := &LoadBalancerSelect{LoadBalancerQuery: lbq}
	sbuild.label = loadbalancer.Label
	sbuild.flds, sbuild.scan = &lbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoadBalancerSelect configured with the given aggregations.
func (lbq *LoadBalancerQuery) Aggregate(fns ...AggregateFunc) *LoadBalancerSelect {
	return lbq.Select().Aggregate(fns...)
}

func (lbq *LoadBalancerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lbq); err != nil {
				return err
			}
		}
	}
	for _, f := range lbq.ctx.Fields {
		if !loadbalancer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lbq.path != nil {
		prev, err := lbq.path(ctx)
		if err != nil {
			return err
		}
		lbq.sql = prev
	}
	return nil
}

func (lbq *LoadBalancerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoadBalancer, error) {
	var (
		nodes       = []*LoadBalancer{}
		_spec       = lbq.querySpec()
		loadedTypes = [1]bool{
			lbq.withPanierLoadBalancer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoadBalancer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoadBalancer{config: lbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lbq.withPanierLoadBalancer; query != nil {
		if err := lbq.loadPanierLoadBalancer(ctx, query, nodes,
			func(n *LoadBalancer) { n.Edges.PanierLoadBalancer = []*PanierLoadBalancer{} },
			func(n *LoadBalancer, e *PanierLoadBalancer) {
				n.Edges.PanierLoadBalancer = append(n.Edges.PanierLoadBalancer, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lbq *LoadBalancerQuery) loadPanierLoadBalancer(ctx context.Context, query *PanierLoadBalancerQuery, nodes []*LoadBalancer, init func(*LoadBalancer), assign func(*LoadBalancer, *PanierLoadBalancer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*LoadBalancer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PanierLoadBalancer(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(loadbalancer.PanierLoadBalancerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.loadBalancer_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "loadBalancer_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "loadBalancer_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (lbq *LoadBalancerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lbq.querySpec()
	_spec.Node.Columns = lbq.ctx.Fields
	if len(lbq.ctx.Fields) > 0 {
		_spec.Unique = lbq.ctx.Unique != nil && *lbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lbq.driver, _spec)
}

func (lbq *LoadBalancerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loadbalancer.Table, loadbalancer.Columns, sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeInt))
	_spec.From = lbq.sql
	if unique := lbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lbq.path != nil {
		_spec.Unique = true
	}
	if fields := lbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loadbalancer.FieldID)
		for i := range fields {
			if fields[i] != loadbalancer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lbq *LoadBalancerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lbq.driver.Dialect())
	t1 := builder.Table(loadbalancer.Table)
	columns := lbq.ctx.Fields
	if len(columns) == 0 {
		columns = loadbalancer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lbq.sql != nil {
		selector = lbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lbq.ctx.Unique != nil && *lbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lbq.predicates {
		p(selector)
	}
	for _, p := range lbq.order {
		p(selector)
	}
	if offset := lbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoadBalancerGroupBy is the group-by builder for LoadBalancer entities.
type LoadBalancerGroupBy struct {
	selector
	build *LoadBalancerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lbgb *LoadBalancerGroupBy) Aggregate(fns ...AggregateFunc) *LoadBalancerGroupBy {
	lbgb.fns = append(lbgb.fns, fns...)
	return lbgb
}

// Scan applies the selector query and scans the result into the given value.
func (lbgb *LoadBalancerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lbgb.build.ctx, "GroupBy")
	if err := lbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadBalancerQuery, *LoadBalancerGroupBy](ctx, lbgb.build, lbgb, lbgb.build.inters, v)
}

func (lbgb *LoadBalancerGroupBy) sqlScan(ctx context.Context, root *LoadBalancerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lbgb.fns))
	for _, fn := range lbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lbgb.flds)+len(lbgb.fns))
		for _, f := range *lbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoadBalancerSelect is the builder for selecting fields of LoadBalancer entities.
type LoadBalancerSelect struct {
	*LoadBalancerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lbs *LoadBalancerSelect) Aggregate(fns ...AggregateFunc) *LoadBalancerSelect {
	lbs.fns = append(lbs.fns, fns...)
	return lbs
}

// Scan applies the selector query and scans the result into the given value.
func (lbs *LoadBalancerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lbs.ctx, "Select")
	if err := lbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadBalancerQuery, *LoadBalancerSelect](ctx, lbs.LoadBalancerQuery, lbs, lbs.inters, v)
}

func (lbs *LoadBalancerSelect) sqlScan(ctx context.Context, root *LoadBalancerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lbs.fns))
	for _, fn := range lbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}