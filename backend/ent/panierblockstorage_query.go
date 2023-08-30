// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/blockstorage"
	"backend/ent/panier"
	"backend/ent/panierblockstorage"
	"backend/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierBlockStorageQuery is the builder for querying PanierBlockStorage entities.
type PanierBlockStorageQuery struct {
	config
	ctx              *QueryContext
	order            []panierblockstorage.OrderOption
	inters           []Interceptor
	predicates       []predicate.PanierBlockStorage
	withPanier       *PanierQuery
	withBlockStorage *BlockStorageQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PanierBlockStorageQuery builder.
func (pbsq *PanierBlockStorageQuery) Where(ps ...predicate.PanierBlockStorage) *PanierBlockStorageQuery {
	pbsq.predicates = append(pbsq.predicates, ps...)
	return pbsq
}

// Limit the number of records to be returned by this query.
func (pbsq *PanierBlockStorageQuery) Limit(limit int) *PanierBlockStorageQuery {
	pbsq.ctx.Limit = &limit
	return pbsq
}

// Offset to start from.
func (pbsq *PanierBlockStorageQuery) Offset(offset int) *PanierBlockStorageQuery {
	pbsq.ctx.Offset = &offset
	return pbsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pbsq *PanierBlockStorageQuery) Unique(unique bool) *PanierBlockStorageQuery {
	pbsq.ctx.Unique = &unique
	return pbsq
}

// Order specifies how the records should be ordered.
func (pbsq *PanierBlockStorageQuery) Order(o ...panierblockstorage.OrderOption) *PanierBlockStorageQuery {
	pbsq.order = append(pbsq.order, o...)
	return pbsq
}

// QueryPanier chains the current query on the "panier" edge.
func (pbsq *PanierBlockStorageQuery) QueryPanier() *PanierQuery {
	query := (&PanierClient{config: pbsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pbsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pbsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panierblockstorage.Table, panierblockstorage.FieldID, selector),
			sqlgraph.To(panier.Table, panier.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panierblockstorage.PanierTable, panierblockstorage.PanierColumn),
		)
		fromU = sqlgraph.SetNeighbors(pbsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBlockStorage chains the current query on the "blockStorage" edge.
func (pbsq *PanierBlockStorageQuery) QueryBlockStorage() *BlockStorageQuery {
	query := (&BlockStorageClient{config: pbsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pbsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pbsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panierblockstorage.Table, panierblockstorage.FieldID, selector),
			sqlgraph.To(blockstorage.Table, blockstorage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panierblockstorage.BlockStorageTable, panierblockstorage.BlockStorageColumn),
		)
		fromU = sqlgraph.SetNeighbors(pbsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PanierBlockStorage entity from the query.
// Returns a *NotFoundError when no PanierBlockStorage was found.
func (pbsq *PanierBlockStorageQuery) First(ctx context.Context) (*PanierBlockStorage, error) {
	nodes, err := pbsq.Limit(1).All(setContextOp(ctx, pbsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{panierblockstorage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) FirstX(ctx context.Context) *PanierBlockStorage {
	node, err := pbsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PanierBlockStorage ID from the query.
// Returns a *NotFoundError when no PanierBlockStorage ID was found.
func (pbsq *PanierBlockStorageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pbsq.Limit(1).IDs(setContextOp(ctx, pbsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{panierblockstorage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) FirstIDX(ctx context.Context) int {
	id, err := pbsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PanierBlockStorage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PanierBlockStorage entity is found.
// Returns a *NotFoundError when no PanierBlockStorage entities are found.
func (pbsq *PanierBlockStorageQuery) Only(ctx context.Context) (*PanierBlockStorage, error) {
	nodes, err := pbsq.Limit(2).All(setContextOp(ctx, pbsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{panierblockstorage.Label}
	default:
		return nil, &NotSingularError{panierblockstorage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) OnlyX(ctx context.Context) *PanierBlockStorage {
	node, err := pbsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PanierBlockStorage ID in the query.
// Returns a *NotSingularError when more than one PanierBlockStorage ID is found.
// Returns a *NotFoundError when no entities are found.
func (pbsq *PanierBlockStorageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pbsq.Limit(2).IDs(setContextOp(ctx, pbsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{panierblockstorage.Label}
	default:
		err = &NotSingularError{panierblockstorage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) OnlyIDX(ctx context.Context) int {
	id, err := pbsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PanierBlockStorages.
func (pbsq *PanierBlockStorageQuery) All(ctx context.Context) ([]*PanierBlockStorage, error) {
	ctx = setContextOp(ctx, pbsq.ctx, "All")
	if err := pbsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PanierBlockStorage, *PanierBlockStorageQuery]()
	return withInterceptors[[]*PanierBlockStorage](ctx, pbsq, qr, pbsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) AllX(ctx context.Context) []*PanierBlockStorage {
	nodes, err := pbsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PanierBlockStorage IDs.
func (pbsq *PanierBlockStorageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pbsq.ctx.Unique == nil && pbsq.path != nil {
		pbsq.Unique(true)
	}
	ctx = setContextOp(ctx, pbsq.ctx, "IDs")
	if err = pbsq.Select(panierblockstorage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) IDsX(ctx context.Context) []int {
	ids, err := pbsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pbsq *PanierBlockStorageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pbsq.ctx, "Count")
	if err := pbsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pbsq, querierCount[*PanierBlockStorageQuery](), pbsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) CountX(ctx context.Context) int {
	count, err := pbsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pbsq *PanierBlockStorageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pbsq.ctx, "Exist")
	switch _, err := pbsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pbsq *PanierBlockStorageQuery) ExistX(ctx context.Context) bool {
	exist, err := pbsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PanierBlockStorageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pbsq *PanierBlockStorageQuery) Clone() *PanierBlockStorageQuery {
	if pbsq == nil {
		return nil
	}
	return &PanierBlockStorageQuery{
		config:           pbsq.config,
		ctx:              pbsq.ctx.Clone(),
		order:            append([]panierblockstorage.OrderOption{}, pbsq.order...),
		inters:           append([]Interceptor{}, pbsq.inters...),
		predicates:       append([]predicate.PanierBlockStorage{}, pbsq.predicates...),
		withPanier:       pbsq.withPanier.Clone(),
		withBlockStorage: pbsq.withBlockStorage.Clone(),
		// clone intermediate query.
		sql:  pbsq.sql.Clone(),
		path: pbsq.path,
	}
}

// WithPanier tells the query-builder to eager-load the nodes that are connected to
// the "panier" edge. The optional arguments are used to configure the query builder of the edge.
func (pbsq *PanierBlockStorageQuery) WithPanier(opts ...func(*PanierQuery)) *PanierBlockStorageQuery {
	query := (&PanierClient{config: pbsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pbsq.withPanier = query
	return pbsq
}

// WithBlockStorage tells the query-builder to eager-load the nodes that are connected to
// the "blockStorage" edge. The optional arguments are used to configure the query builder of the edge.
func (pbsq *PanierBlockStorageQuery) WithBlockStorage(opts ...func(*BlockStorageQuery)) *PanierBlockStorageQuery {
	query := (&BlockStorageClient{config: pbsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pbsq.withBlockStorage = query
	return pbsq
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
//	client.PanierBlockStorage.Query().
//		GroupBy(panierblockstorage.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pbsq *PanierBlockStorageQuery) GroupBy(field string, fields ...string) *PanierBlockStorageGroupBy {
	pbsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PanierBlockStorageGroupBy{build: pbsq}
	grbuild.flds = &pbsq.ctx.Fields
	grbuild.label = panierblockstorage.Label
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
//	client.PanierBlockStorage.Query().
//		Select(panierblockstorage.FieldQuantity).
//		Scan(ctx, &v)
func (pbsq *PanierBlockStorageQuery) Select(fields ...string) *PanierBlockStorageSelect {
	pbsq.ctx.Fields = append(pbsq.ctx.Fields, fields...)
	sbuild := &PanierBlockStorageSelect{PanierBlockStorageQuery: pbsq}
	sbuild.label = panierblockstorage.Label
	sbuild.flds, sbuild.scan = &pbsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PanierBlockStorageSelect configured with the given aggregations.
func (pbsq *PanierBlockStorageQuery) Aggregate(fns ...AggregateFunc) *PanierBlockStorageSelect {
	return pbsq.Select().Aggregate(fns...)
}

func (pbsq *PanierBlockStorageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pbsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pbsq); err != nil {
				return err
			}
		}
	}
	for _, f := range pbsq.ctx.Fields {
		if !panierblockstorage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pbsq.path != nil {
		prev, err := pbsq.path(ctx)
		if err != nil {
			return err
		}
		pbsq.sql = prev
	}
	return nil
}

func (pbsq *PanierBlockStorageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PanierBlockStorage, error) {
	var (
		nodes       = []*PanierBlockStorage{}
		withFKs     = pbsq.withFKs
		_spec       = pbsq.querySpec()
		loadedTypes = [2]bool{
			pbsq.withPanier != nil,
			pbsq.withBlockStorage != nil,
		}
	)
	if pbsq.withPanier != nil || pbsq.withBlockStorage != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, panierblockstorage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PanierBlockStorage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PanierBlockStorage{config: pbsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pbsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pbsq.withPanier; query != nil {
		if err := pbsq.loadPanier(ctx, query, nodes, nil,
			func(n *PanierBlockStorage, e *Panier) { n.Edges.Panier = e }); err != nil {
			return nil, err
		}
	}
	if query := pbsq.withBlockStorage; query != nil {
		if err := pbsq.loadBlockStorage(ctx, query, nodes, nil,
			func(n *PanierBlockStorage, e *BlockStorage) { n.Edges.BlockStorage = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pbsq *PanierBlockStorageQuery) loadPanier(ctx context.Context, query *PanierQuery, nodes []*PanierBlockStorage, init func(*PanierBlockStorage), assign func(*PanierBlockStorage, *Panier)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierBlockStorage)
	for i := range nodes {
		if nodes[i].panier_panier_block_storage == nil {
			continue
		}
		fk := *nodes[i].panier_panier_block_storage
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
			return fmt.Errorf(`unexpected foreign-key "panier_panier_block_storage" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pbsq *PanierBlockStorageQuery) loadBlockStorage(ctx context.Context, query *BlockStorageQuery, nodes []*PanierBlockStorage, init func(*PanierBlockStorage), assign func(*PanierBlockStorage, *BlockStorage)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierBlockStorage)
	for i := range nodes {
		if nodes[i].block_storage_panier_block_storage == nil {
			continue
		}
		fk := *nodes[i].block_storage_panier_block_storage
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(blockstorage.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "block_storage_panier_block_storage" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pbsq *PanierBlockStorageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pbsq.querySpec()
	_spec.Node.Columns = pbsq.ctx.Fields
	if len(pbsq.ctx.Fields) > 0 {
		_spec.Unique = pbsq.ctx.Unique != nil && *pbsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pbsq.driver, _spec)
}

func (pbsq *PanierBlockStorageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(panierblockstorage.Table, panierblockstorage.Columns, sqlgraph.NewFieldSpec(panierblockstorage.FieldID, field.TypeInt))
	_spec.From = pbsq.sql
	if unique := pbsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pbsq.path != nil {
		_spec.Unique = true
	}
	if fields := pbsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, panierblockstorage.FieldID)
		for i := range fields {
			if fields[i] != panierblockstorage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pbsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pbsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pbsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pbsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pbsq *PanierBlockStorageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pbsq.driver.Dialect())
	t1 := builder.Table(panierblockstorage.Table)
	columns := pbsq.ctx.Fields
	if len(columns) == 0 {
		columns = panierblockstorage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pbsq.sql != nil {
		selector = pbsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pbsq.ctx.Unique != nil && *pbsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pbsq.predicates {
		p(selector)
	}
	for _, p := range pbsq.order {
		p(selector)
	}
	if offset := pbsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pbsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PanierBlockStorageGroupBy is the group-by builder for PanierBlockStorage entities.
type PanierBlockStorageGroupBy struct {
	selector
	build *PanierBlockStorageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pbsgb *PanierBlockStorageGroupBy) Aggregate(fns ...AggregateFunc) *PanierBlockStorageGroupBy {
	pbsgb.fns = append(pbsgb.fns, fns...)
	return pbsgb
}

// Scan applies the selector query and scans the result into the given value.
func (pbsgb *PanierBlockStorageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pbsgb.build.ctx, "GroupBy")
	if err := pbsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierBlockStorageQuery, *PanierBlockStorageGroupBy](ctx, pbsgb.build, pbsgb, pbsgb.build.inters, v)
}

func (pbsgb *PanierBlockStorageGroupBy) sqlScan(ctx context.Context, root *PanierBlockStorageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pbsgb.fns))
	for _, fn := range pbsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pbsgb.flds)+len(pbsgb.fns))
		for _, f := range *pbsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pbsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pbsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PanierBlockStorageSelect is the builder for selecting fields of PanierBlockStorage entities.
type PanierBlockStorageSelect struct {
	*PanierBlockStorageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pbss *PanierBlockStorageSelect) Aggregate(fns ...AggregateFunc) *PanierBlockStorageSelect {
	pbss.fns = append(pbss.fns, fns...)
	return pbss
}

// Scan applies the selector query and scans the result into the given value.
func (pbss *PanierBlockStorageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pbss.ctx, "Select")
	if err := pbss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierBlockStorageQuery, *PanierBlockStorageSelect](ctx, pbss.PanierBlockStorageQuery, pbss, pbss.inters, v)
}

func (pbss *PanierBlockStorageSelect) sqlScan(ctx context.Context, root *PanierBlockStorageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pbss.fns))
	for _, fn := range pbss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pbss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pbss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
