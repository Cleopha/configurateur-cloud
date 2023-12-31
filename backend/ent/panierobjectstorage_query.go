// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/objectstorage"
	"backend/ent/panier"
	"backend/ent/panierobjectstorage"
	"backend/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierObjectStorageQuery is the builder for querying PanierObjectStorage entities.
type PanierObjectStorageQuery struct {
	config
	ctx               *QueryContext
	order             []panierobjectstorage.OrderOption
	inters            []Interceptor
	predicates        []predicate.PanierObjectStorage
	withPanier        *PanierQuery
	withObjectStorage *ObjectStorageQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PanierObjectStorageQuery builder.
func (posq *PanierObjectStorageQuery) Where(ps ...predicate.PanierObjectStorage) *PanierObjectStorageQuery {
	posq.predicates = append(posq.predicates, ps...)
	return posq
}

// Limit the number of records to be returned by this query.
func (posq *PanierObjectStorageQuery) Limit(limit int) *PanierObjectStorageQuery {
	posq.ctx.Limit = &limit
	return posq
}

// Offset to start from.
func (posq *PanierObjectStorageQuery) Offset(offset int) *PanierObjectStorageQuery {
	posq.ctx.Offset = &offset
	return posq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (posq *PanierObjectStorageQuery) Unique(unique bool) *PanierObjectStorageQuery {
	posq.ctx.Unique = &unique
	return posq
}

// Order specifies how the records should be ordered.
func (posq *PanierObjectStorageQuery) Order(o ...panierobjectstorage.OrderOption) *PanierObjectStorageQuery {
	posq.order = append(posq.order, o...)
	return posq
}

// QueryPanier chains the current query on the "panier" edge.
func (posq *PanierObjectStorageQuery) QueryPanier() *PanierQuery {
	query := (&PanierClient{config: posq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := posq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := posq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panierobjectstorage.Table, panierobjectstorage.FieldID, selector),
			sqlgraph.To(panier.Table, panier.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panierobjectstorage.PanierTable, panierobjectstorage.PanierColumn),
		)
		fromU = sqlgraph.SetNeighbors(posq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryObjectStorage chains the current query on the "objectStorage" edge.
func (posq *PanierObjectStorageQuery) QueryObjectStorage() *ObjectStorageQuery {
	query := (&ObjectStorageClient{config: posq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := posq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := posq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panierobjectstorage.Table, panierobjectstorage.FieldID, selector),
			sqlgraph.To(objectstorage.Table, objectstorage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panierobjectstorage.ObjectStorageTable, panierobjectstorage.ObjectStorageColumn),
		)
		fromU = sqlgraph.SetNeighbors(posq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PanierObjectStorage entity from the query.
// Returns a *NotFoundError when no PanierObjectStorage was found.
func (posq *PanierObjectStorageQuery) First(ctx context.Context) (*PanierObjectStorage, error) {
	nodes, err := posq.Limit(1).All(setContextOp(ctx, posq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{panierobjectstorage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) FirstX(ctx context.Context) *PanierObjectStorage {
	node, err := posq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PanierObjectStorage ID from the query.
// Returns a *NotFoundError when no PanierObjectStorage ID was found.
func (posq *PanierObjectStorageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = posq.Limit(1).IDs(setContextOp(ctx, posq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{panierobjectstorage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) FirstIDX(ctx context.Context) int {
	id, err := posq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PanierObjectStorage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PanierObjectStorage entity is found.
// Returns a *NotFoundError when no PanierObjectStorage entities are found.
func (posq *PanierObjectStorageQuery) Only(ctx context.Context) (*PanierObjectStorage, error) {
	nodes, err := posq.Limit(2).All(setContextOp(ctx, posq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{panierobjectstorage.Label}
	default:
		return nil, &NotSingularError{panierobjectstorage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) OnlyX(ctx context.Context) *PanierObjectStorage {
	node, err := posq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PanierObjectStorage ID in the query.
// Returns a *NotSingularError when more than one PanierObjectStorage ID is found.
// Returns a *NotFoundError when no entities are found.
func (posq *PanierObjectStorageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = posq.Limit(2).IDs(setContextOp(ctx, posq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{panierobjectstorage.Label}
	default:
		err = &NotSingularError{panierobjectstorage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) OnlyIDX(ctx context.Context) int {
	id, err := posq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PanierObjectStorages.
func (posq *PanierObjectStorageQuery) All(ctx context.Context) ([]*PanierObjectStorage, error) {
	ctx = setContextOp(ctx, posq.ctx, "All")
	if err := posq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PanierObjectStorage, *PanierObjectStorageQuery]()
	return withInterceptors[[]*PanierObjectStorage](ctx, posq, qr, posq.inters)
}

// AllX is like All, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) AllX(ctx context.Context) []*PanierObjectStorage {
	nodes, err := posq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PanierObjectStorage IDs.
func (posq *PanierObjectStorageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if posq.ctx.Unique == nil && posq.path != nil {
		posq.Unique(true)
	}
	ctx = setContextOp(ctx, posq.ctx, "IDs")
	if err = posq.Select(panierobjectstorage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) IDsX(ctx context.Context) []int {
	ids, err := posq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (posq *PanierObjectStorageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, posq.ctx, "Count")
	if err := posq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, posq, querierCount[*PanierObjectStorageQuery](), posq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) CountX(ctx context.Context) int {
	count, err := posq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (posq *PanierObjectStorageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, posq.ctx, "Exist")
	switch _, err := posq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (posq *PanierObjectStorageQuery) ExistX(ctx context.Context) bool {
	exist, err := posq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PanierObjectStorageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (posq *PanierObjectStorageQuery) Clone() *PanierObjectStorageQuery {
	if posq == nil {
		return nil
	}
	return &PanierObjectStorageQuery{
		config:            posq.config,
		ctx:               posq.ctx.Clone(),
		order:             append([]panierobjectstorage.OrderOption{}, posq.order...),
		inters:            append([]Interceptor{}, posq.inters...),
		predicates:        append([]predicate.PanierObjectStorage{}, posq.predicates...),
		withPanier:        posq.withPanier.Clone(),
		withObjectStorage: posq.withObjectStorage.Clone(),
		// clone intermediate query.
		sql:  posq.sql.Clone(),
		path: posq.path,
	}
}

// WithPanier tells the query-builder to eager-load the nodes that are connected to
// the "panier" edge. The optional arguments are used to configure the query builder of the edge.
func (posq *PanierObjectStorageQuery) WithPanier(opts ...func(*PanierQuery)) *PanierObjectStorageQuery {
	query := (&PanierClient{config: posq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	posq.withPanier = query
	return posq
}

// WithObjectStorage tells the query-builder to eager-load the nodes that are connected to
// the "objectStorage" edge. The optional arguments are used to configure the query builder of the edge.
func (posq *PanierObjectStorageQuery) WithObjectStorage(opts ...func(*ObjectStorageQuery)) *PanierObjectStorageQuery {
	query := (&ObjectStorageClient{config: posq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	posq.withObjectStorage = query
	return posq
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
//	client.PanierObjectStorage.Query().
//		GroupBy(panierobjectstorage.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (posq *PanierObjectStorageQuery) GroupBy(field string, fields ...string) *PanierObjectStorageGroupBy {
	posq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PanierObjectStorageGroupBy{build: posq}
	grbuild.flds = &posq.ctx.Fields
	grbuild.label = panierobjectstorage.Label
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
//	client.PanierObjectStorage.Query().
//		Select(panierobjectstorage.FieldQuantity).
//		Scan(ctx, &v)
func (posq *PanierObjectStorageQuery) Select(fields ...string) *PanierObjectStorageSelect {
	posq.ctx.Fields = append(posq.ctx.Fields, fields...)
	sbuild := &PanierObjectStorageSelect{PanierObjectStorageQuery: posq}
	sbuild.label = panierobjectstorage.Label
	sbuild.flds, sbuild.scan = &posq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PanierObjectStorageSelect configured with the given aggregations.
func (posq *PanierObjectStorageQuery) Aggregate(fns ...AggregateFunc) *PanierObjectStorageSelect {
	return posq.Select().Aggregate(fns...)
}

func (posq *PanierObjectStorageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range posq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, posq); err != nil {
				return err
			}
		}
	}
	for _, f := range posq.ctx.Fields {
		if !panierobjectstorage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if posq.path != nil {
		prev, err := posq.path(ctx)
		if err != nil {
			return err
		}
		posq.sql = prev
	}
	return nil
}

func (posq *PanierObjectStorageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PanierObjectStorage, error) {
	var (
		nodes       = []*PanierObjectStorage{}
		withFKs     = posq.withFKs
		_spec       = posq.querySpec()
		loadedTypes = [2]bool{
			posq.withPanier != nil,
			posq.withObjectStorage != nil,
		}
	)
	if posq.withPanier != nil || posq.withObjectStorage != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, panierobjectstorage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PanierObjectStorage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PanierObjectStorage{config: posq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, posq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := posq.withPanier; query != nil {
		if err := posq.loadPanier(ctx, query, nodes, nil,
			func(n *PanierObjectStorage, e *Panier) { n.Edges.Panier = e }); err != nil {
			return nil, err
		}
	}
	if query := posq.withObjectStorage; query != nil {
		if err := posq.loadObjectStorage(ctx, query, nodes, nil,
			func(n *PanierObjectStorage, e *ObjectStorage) { n.Edges.ObjectStorage = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (posq *PanierObjectStorageQuery) loadPanier(ctx context.Context, query *PanierQuery, nodes []*PanierObjectStorage, init func(*PanierObjectStorage), assign func(*PanierObjectStorage, *Panier)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierObjectStorage)
	for i := range nodes {
		if nodes[i].panier_panier_object_storage == nil {
			continue
		}
		fk := *nodes[i].panier_panier_object_storage
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
			return fmt.Errorf(`unexpected foreign-key "panier_panier_object_storage" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (posq *PanierObjectStorageQuery) loadObjectStorage(ctx context.Context, query *ObjectStorageQuery, nodes []*PanierObjectStorage, init func(*PanierObjectStorage), assign func(*PanierObjectStorage, *ObjectStorage)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierObjectStorage)
	for i := range nodes {
		if nodes[i].objectStorage_id == nil {
			continue
		}
		fk := *nodes[i].objectStorage_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(objectstorage.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "objectStorage_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (posq *PanierObjectStorageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := posq.querySpec()
	_spec.Node.Columns = posq.ctx.Fields
	if len(posq.ctx.Fields) > 0 {
		_spec.Unique = posq.ctx.Unique != nil && *posq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, posq.driver, _spec)
}

func (posq *PanierObjectStorageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(panierobjectstorage.Table, panierobjectstorage.Columns, sqlgraph.NewFieldSpec(panierobjectstorage.FieldID, field.TypeInt))
	_spec.From = posq.sql
	if unique := posq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if posq.path != nil {
		_spec.Unique = true
	}
	if fields := posq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, panierobjectstorage.FieldID)
		for i := range fields {
			if fields[i] != panierobjectstorage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := posq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := posq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := posq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := posq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (posq *PanierObjectStorageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(posq.driver.Dialect())
	t1 := builder.Table(panierobjectstorage.Table)
	columns := posq.ctx.Fields
	if len(columns) == 0 {
		columns = panierobjectstorage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if posq.sql != nil {
		selector = posq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if posq.ctx.Unique != nil && *posq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range posq.predicates {
		p(selector)
	}
	for _, p := range posq.order {
		p(selector)
	}
	if offset := posq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := posq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PanierObjectStorageGroupBy is the group-by builder for PanierObjectStorage entities.
type PanierObjectStorageGroupBy struct {
	selector
	build *PanierObjectStorageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (posgb *PanierObjectStorageGroupBy) Aggregate(fns ...AggregateFunc) *PanierObjectStorageGroupBy {
	posgb.fns = append(posgb.fns, fns...)
	return posgb
}

// Scan applies the selector query and scans the result into the given value.
func (posgb *PanierObjectStorageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, posgb.build.ctx, "GroupBy")
	if err := posgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierObjectStorageQuery, *PanierObjectStorageGroupBy](ctx, posgb.build, posgb, posgb.build.inters, v)
}

func (posgb *PanierObjectStorageGroupBy) sqlScan(ctx context.Context, root *PanierObjectStorageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(posgb.fns))
	for _, fn := range posgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*posgb.flds)+len(posgb.fns))
		for _, f := range *posgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*posgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := posgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PanierObjectStorageSelect is the builder for selecting fields of PanierObjectStorage entities.
type PanierObjectStorageSelect struct {
	*PanierObjectStorageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (poss *PanierObjectStorageSelect) Aggregate(fns ...AggregateFunc) *PanierObjectStorageSelect {
	poss.fns = append(poss.fns, fns...)
	return poss
}

// Scan applies the selector query and scans the result into the given value.
func (poss *PanierObjectStorageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, poss.ctx, "Select")
	if err := poss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierObjectStorageQuery, *PanierObjectStorageSelect](ctx, poss.PanierObjectStorageQuery, poss, poss.inters, v)
}

func (poss *PanierObjectStorageSelect) sqlScan(ctx context.Context, root *PanierObjectStorageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(poss.fns))
	for _, fn := range poss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*poss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := poss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
