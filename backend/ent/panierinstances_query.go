// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/instances"
	"backend/ent/panier"
	"backend/ent/panierinstances"
	"backend/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierInstancesQuery is the builder for querying PanierInstances entities.
type PanierInstancesQuery struct {
	config
	ctx          *QueryContext
	order        []panierinstances.OrderOption
	inters       []Interceptor
	predicates   []predicate.PanierInstances
	withPanier   *PanierQuery
	withInstance *InstancesQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PanierInstancesQuery builder.
func (piq *PanierInstancesQuery) Where(ps ...predicate.PanierInstances) *PanierInstancesQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit the number of records to be returned by this query.
func (piq *PanierInstancesQuery) Limit(limit int) *PanierInstancesQuery {
	piq.ctx.Limit = &limit
	return piq
}

// Offset to start from.
func (piq *PanierInstancesQuery) Offset(offset int) *PanierInstancesQuery {
	piq.ctx.Offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *PanierInstancesQuery) Unique(unique bool) *PanierInstancesQuery {
	piq.ctx.Unique = &unique
	return piq
}

// Order specifies how the records should be ordered.
func (piq *PanierInstancesQuery) Order(o ...panierinstances.OrderOption) *PanierInstancesQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// QueryPanier chains the current query on the "panier" edge.
func (piq *PanierInstancesQuery) QueryPanier() *PanierQuery {
	query := (&PanierClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panierinstances.Table, panierinstances.FieldID, selector),
			sqlgraph.To(panier.Table, panier.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panierinstances.PanierTable, panierinstances.PanierColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstance chains the current query on the "instance" edge.
func (piq *PanierInstancesQuery) QueryInstance() *InstancesQuery {
	query := (&InstancesClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panierinstances.Table, panierinstances.FieldID, selector),
			sqlgraph.To(instances.Table, instances.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panierinstances.InstanceTable, panierinstances.InstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PanierInstances entity from the query.
// Returns a *NotFoundError when no PanierInstances was found.
func (piq *PanierInstancesQuery) First(ctx context.Context) (*PanierInstances, error) {
	nodes, err := piq.Limit(1).All(setContextOp(ctx, piq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{panierinstances.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *PanierInstancesQuery) FirstX(ctx context.Context) *PanierInstances {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PanierInstances ID from the query.
// Returns a *NotFoundError when no PanierInstances ID was found.
func (piq *PanierInstancesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(1).IDs(setContextOp(ctx, piq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{panierinstances.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *PanierInstancesQuery) FirstIDX(ctx context.Context) int {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PanierInstances entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PanierInstances entity is found.
// Returns a *NotFoundError when no PanierInstances entities are found.
func (piq *PanierInstancesQuery) Only(ctx context.Context) (*PanierInstances, error) {
	nodes, err := piq.Limit(2).All(setContextOp(ctx, piq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{panierinstances.Label}
	default:
		return nil, &NotSingularError{panierinstances.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *PanierInstancesQuery) OnlyX(ctx context.Context) *PanierInstances {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PanierInstances ID in the query.
// Returns a *NotSingularError when more than one PanierInstances ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *PanierInstancesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(2).IDs(setContextOp(ctx, piq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{panierinstances.Label}
	default:
		err = &NotSingularError{panierinstances.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *PanierInstancesQuery) OnlyIDX(ctx context.Context) int {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PanierInstancesSlice.
func (piq *PanierInstancesQuery) All(ctx context.Context) ([]*PanierInstances, error) {
	ctx = setContextOp(ctx, piq.ctx, "All")
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PanierInstances, *PanierInstancesQuery]()
	return withInterceptors[[]*PanierInstances](ctx, piq, qr, piq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piq *PanierInstancesQuery) AllX(ctx context.Context) []*PanierInstances {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PanierInstances IDs.
func (piq *PanierInstancesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if piq.ctx.Unique == nil && piq.path != nil {
		piq.Unique(true)
	}
	ctx = setContextOp(ctx, piq.ctx, "IDs")
	if err = piq.Select(panierinstances.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *PanierInstancesQuery) IDsX(ctx context.Context) []int {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *PanierInstancesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piq.ctx, "Count")
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piq, querierCount[*PanierInstancesQuery](), piq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piq *PanierInstancesQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *PanierInstancesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piq.ctx, "Exist")
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *PanierInstancesQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PanierInstancesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *PanierInstancesQuery) Clone() *PanierInstancesQuery {
	if piq == nil {
		return nil
	}
	return &PanierInstancesQuery{
		config:       piq.config,
		ctx:          piq.ctx.Clone(),
		order:        append([]panierinstances.OrderOption{}, piq.order...),
		inters:       append([]Interceptor{}, piq.inters...),
		predicates:   append([]predicate.PanierInstances{}, piq.predicates...),
		withPanier:   piq.withPanier.Clone(),
		withInstance: piq.withInstance.Clone(),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
	}
}

// WithPanier tells the query-builder to eager-load the nodes that are connected to
// the "panier" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *PanierInstancesQuery) WithPanier(opts ...func(*PanierQuery)) *PanierInstancesQuery {
	query := (&PanierClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withPanier = query
	return piq
}

// WithInstance tells the query-builder to eager-load the nodes that are connected to
// the "instance" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *PanierInstancesQuery) WithInstance(opts ...func(*InstancesQuery)) *PanierInstancesQuery {
	query := (&InstancesClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withInstance = query
	return piq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PanierInstances.Query().
//		GroupBy(panierinstances.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *PanierInstancesQuery) GroupBy(field string, fields ...string) *PanierInstancesGroupBy {
	piq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PanierInstancesGroupBy{build: piq}
	grbuild.flds = &piq.ctx.Fields
	grbuild.label = panierinstances.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//	}
//
//	client.PanierInstances.Query().
//		Select(panierinstances.FieldQuantity).
//		Scan(ctx, &v)
func (piq *PanierInstancesQuery) Select(fields ...string) *PanierInstancesSelect {
	piq.ctx.Fields = append(piq.ctx.Fields, fields...)
	sbuild := &PanierInstancesSelect{PanierInstancesQuery: piq}
	sbuild.label = panierinstances.Label
	sbuild.flds, sbuild.scan = &piq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PanierInstancesSelect configured with the given aggregations.
func (piq *PanierInstancesQuery) Aggregate(fns ...AggregateFunc) *PanierInstancesSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *PanierInstancesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piq); err != nil {
				return err
			}
		}
	}
	for _, f := range piq.ctx.Fields {
		if !panierinstances.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *PanierInstancesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PanierInstances, error) {
	var (
		nodes       = []*PanierInstances{}
		withFKs     = piq.withFKs
		_spec       = piq.querySpec()
		loadedTypes = [2]bool{
			piq.withPanier != nil,
			piq.withInstance != nil,
		}
	)
	if piq.withPanier != nil || piq.withInstance != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, panierinstances.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PanierInstances).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PanierInstances{config: piq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := piq.withPanier; query != nil {
		if err := piq.loadPanier(ctx, query, nodes, nil,
			func(n *PanierInstances, e *Panier) { n.Edges.Panier = e }); err != nil {
			return nil, err
		}
	}
	if query := piq.withInstance; query != nil {
		if err := piq.loadInstance(ctx, query, nodes, nil,
			func(n *PanierInstances, e *Instances) { n.Edges.Instance = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (piq *PanierInstancesQuery) loadPanier(ctx context.Context, query *PanierQuery, nodes []*PanierInstances, init func(*PanierInstances), assign func(*PanierInstances, *Panier)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierInstances)
	for i := range nodes {
		if nodes[i].panier_panier_instances == nil {
			continue
		}
		fk := *nodes[i].panier_panier_instances
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(panier.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "panier_panier_instances" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (piq *PanierInstancesQuery) loadInstance(ctx context.Context, query *InstancesQuery, nodes []*PanierInstances, init func(*PanierInstances), assign func(*PanierInstances, *Instances)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierInstances)
	for i := range nodes {
		if nodes[i].instance_id == nil {
			continue
		}
		fk := *nodes[i].instance_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(instances.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "instance_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (piq *PanierInstancesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.ctx.Fields
	if len(piq.ctx.Fields) > 0 {
		_spec.Unique = piq.ctx.Unique != nil && *piq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *PanierInstancesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(panierinstances.Table, panierinstances.Columns, sqlgraph.NewFieldSpec(panierinstances.FieldID, field.TypeInt))
	_spec.From = piq.sql
	if unique := piq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piq.path != nil {
		_spec.Unique = true
	}
	if fields := piq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, panierinstances.FieldID)
		for i := range fields {
			if fields[i] != panierinstances.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *PanierInstancesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(panierinstances.Table)
	columns := piq.ctx.Fields
	if len(columns) == 0 {
		columns = panierinstances.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.ctx.Unique != nil && *piq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PanierInstancesGroupBy is the group-by builder for PanierInstances entities.
type PanierInstancesGroupBy struct {
	selector
	build *PanierInstancesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *PanierInstancesGroupBy) Aggregate(fns ...AggregateFunc) *PanierInstancesGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the selector query and scans the result into the given value.
func (pigb *PanierInstancesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pigb.build.ctx, "GroupBy")
	if err := pigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierInstancesQuery, *PanierInstancesGroupBy](ctx, pigb.build, pigb, pigb.build.inters, v)
}

func (pigb *PanierInstancesGroupBy) sqlScan(ctx context.Context, root *PanierInstancesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pigb.flds)+len(pigb.fns))
		for _, f := range *pigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PanierInstancesSelect is the builder for selecting fields of PanierInstances entities.
type PanierInstancesSelect struct {
	*PanierInstancesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *PanierInstancesSelect) Aggregate(fns ...AggregateFunc) *PanierInstancesSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *PanierInstancesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pis.ctx, "Select")
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierInstancesQuery, *PanierInstancesSelect](ctx, pis.PanierInstancesQuery, pis, pis.inters, v)
}

func (pis *PanierInstancesSelect) sqlScan(ctx context.Context, root *PanierInstancesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}