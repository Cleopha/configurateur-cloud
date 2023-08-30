// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/ipaddress"
	"backend/ent/panier"
	"backend/ent/panieripaddress"
	"backend/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PanierIPAddressQuery is the builder for querying PanierIPAddress entities.
type PanierIPAddressQuery struct {
	config
	ctx           *QueryContext
	order         []panieripaddress.OrderOption
	inters        []Interceptor
	predicates    []predicate.PanierIPAddress
	withPanier    *PanierQuery
	withIPAddress *IPAddressQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PanierIPAddressQuery builder.
func (piaq *PanierIPAddressQuery) Where(ps ...predicate.PanierIPAddress) *PanierIPAddressQuery {
	piaq.predicates = append(piaq.predicates, ps...)
	return piaq
}

// Limit the number of records to be returned by this query.
func (piaq *PanierIPAddressQuery) Limit(limit int) *PanierIPAddressQuery {
	piaq.ctx.Limit = &limit
	return piaq
}

// Offset to start from.
func (piaq *PanierIPAddressQuery) Offset(offset int) *PanierIPAddressQuery {
	piaq.ctx.Offset = &offset
	return piaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piaq *PanierIPAddressQuery) Unique(unique bool) *PanierIPAddressQuery {
	piaq.ctx.Unique = &unique
	return piaq
}

// Order specifies how the records should be ordered.
func (piaq *PanierIPAddressQuery) Order(o ...panieripaddress.OrderOption) *PanierIPAddressQuery {
	piaq.order = append(piaq.order, o...)
	return piaq
}

// QueryPanier chains the current query on the "panier" edge.
func (piaq *PanierIPAddressQuery) QueryPanier() *PanierQuery {
	query := (&PanierClient{config: piaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panieripaddress.Table, panieripaddress.FieldID, selector),
			sqlgraph.To(panier.Table, panier.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panieripaddress.PanierTable, panieripaddress.PanierColumn),
		)
		fromU = sqlgraph.SetNeighbors(piaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIPAddress chains the current query on the "IPAddress" edge.
func (piaq *PanierIPAddressQuery) QueryIPAddress() *IPAddressQuery {
	query := (&IPAddressClient{config: piaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(panieripaddress.Table, panieripaddress.FieldID, selector),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, panieripaddress.IPAddressTable, panieripaddress.IPAddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(piaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PanierIPAddress entity from the query.
// Returns a *NotFoundError when no PanierIPAddress was found.
func (piaq *PanierIPAddressQuery) First(ctx context.Context) (*PanierIPAddress, error) {
	nodes, err := piaq.Limit(1).All(setContextOp(ctx, piaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{panieripaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) FirstX(ctx context.Context) *PanierIPAddress {
	node, err := piaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PanierIPAddress ID from the query.
// Returns a *NotFoundError when no PanierIPAddress ID was found.
func (piaq *PanierIPAddressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piaq.Limit(1).IDs(setContextOp(ctx, piaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{panieripaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) FirstIDX(ctx context.Context) int {
	id, err := piaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PanierIPAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PanierIPAddress entity is found.
// Returns a *NotFoundError when no PanierIPAddress entities are found.
func (piaq *PanierIPAddressQuery) Only(ctx context.Context) (*PanierIPAddress, error) {
	nodes, err := piaq.Limit(2).All(setContextOp(ctx, piaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{panieripaddress.Label}
	default:
		return nil, &NotSingularError{panieripaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) OnlyX(ctx context.Context) *PanierIPAddress {
	node, err := piaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PanierIPAddress ID in the query.
// Returns a *NotSingularError when more than one PanierIPAddress ID is found.
// Returns a *NotFoundError when no entities are found.
func (piaq *PanierIPAddressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piaq.Limit(2).IDs(setContextOp(ctx, piaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{panieripaddress.Label}
	default:
		err = &NotSingularError{panieripaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) OnlyIDX(ctx context.Context) int {
	id, err := piaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PanierIPAddresses.
func (piaq *PanierIPAddressQuery) All(ctx context.Context) ([]*PanierIPAddress, error) {
	ctx = setContextOp(ctx, piaq.ctx, "All")
	if err := piaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PanierIPAddress, *PanierIPAddressQuery]()
	return withInterceptors[[]*PanierIPAddress](ctx, piaq, qr, piaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) AllX(ctx context.Context) []*PanierIPAddress {
	nodes, err := piaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PanierIPAddress IDs.
func (piaq *PanierIPAddressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if piaq.ctx.Unique == nil && piaq.path != nil {
		piaq.Unique(true)
	}
	ctx = setContextOp(ctx, piaq.ctx, "IDs")
	if err = piaq.Select(panieripaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) IDsX(ctx context.Context) []int {
	ids, err := piaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piaq *PanierIPAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piaq.ctx, "Count")
	if err := piaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piaq, querierCount[*PanierIPAddressQuery](), piaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) CountX(ctx context.Context) int {
	count, err := piaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piaq *PanierIPAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piaq.ctx, "Exist")
	switch _, err := piaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piaq *PanierIPAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := piaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PanierIPAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piaq *PanierIPAddressQuery) Clone() *PanierIPAddressQuery {
	if piaq == nil {
		return nil
	}
	return &PanierIPAddressQuery{
		config:        piaq.config,
		ctx:           piaq.ctx.Clone(),
		order:         append([]panieripaddress.OrderOption{}, piaq.order...),
		inters:        append([]Interceptor{}, piaq.inters...),
		predicates:    append([]predicate.PanierIPAddress{}, piaq.predicates...),
		withPanier:    piaq.withPanier.Clone(),
		withIPAddress: piaq.withIPAddress.Clone(),
		// clone intermediate query.
		sql:  piaq.sql.Clone(),
		path: piaq.path,
	}
}

// WithPanier tells the query-builder to eager-load the nodes that are connected to
// the "panier" edge. The optional arguments are used to configure the query builder of the edge.
func (piaq *PanierIPAddressQuery) WithPanier(opts ...func(*PanierQuery)) *PanierIPAddressQuery {
	query := (&PanierClient{config: piaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piaq.withPanier = query
	return piaq
}

// WithIPAddress tells the query-builder to eager-load the nodes that are connected to
// the "IPAddress" edge. The optional arguments are used to configure the query builder of the edge.
func (piaq *PanierIPAddressQuery) WithIPAddress(opts ...func(*IPAddressQuery)) *PanierIPAddressQuery {
	query := (&IPAddressClient{config: piaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piaq.withIPAddress = query
	return piaq
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
//	client.PanierIPAddress.Query().
//		GroupBy(panieripaddress.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piaq *PanierIPAddressQuery) GroupBy(field string, fields ...string) *PanierIPAddressGroupBy {
	piaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PanierIPAddressGroupBy{build: piaq}
	grbuild.flds = &piaq.ctx.Fields
	grbuild.label = panieripaddress.Label
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
//	client.PanierIPAddress.Query().
//		Select(panieripaddress.FieldQuantity).
//		Scan(ctx, &v)
func (piaq *PanierIPAddressQuery) Select(fields ...string) *PanierIPAddressSelect {
	piaq.ctx.Fields = append(piaq.ctx.Fields, fields...)
	sbuild := &PanierIPAddressSelect{PanierIPAddressQuery: piaq}
	sbuild.label = panieripaddress.Label
	sbuild.flds, sbuild.scan = &piaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PanierIPAddressSelect configured with the given aggregations.
func (piaq *PanierIPAddressQuery) Aggregate(fns ...AggregateFunc) *PanierIPAddressSelect {
	return piaq.Select().Aggregate(fns...)
}

func (piaq *PanierIPAddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piaq); err != nil {
				return err
			}
		}
	}
	for _, f := range piaq.ctx.Fields {
		if !panieripaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piaq.path != nil {
		prev, err := piaq.path(ctx)
		if err != nil {
			return err
		}
		piaq.sql = prev
	}
	return nil
}

func (piaq *PanierIPAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PanierIPAddress, error) {
	var (
		nodes       = []*PanierIPAddress{}
		withFKs     = piaq.withFKs
		_spec       = piaq.querySpec()
		loadedTypes = [2]bool{
			piaq.withPanier != nil,
			piaq.withIPAddress != nil,
		}
	)
	if piaq.withPanier != nil || piaq.withIPAddress != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, panieripaddress.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PanierIPAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PanierIPAddress{config: piaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := piaq.withPanier; query != nil {
		if err := piaq.loadPanier(ctx, query, nodes, nil,
			func(n *PanierIPAddress, e *Panier) { n.Edges.Panier = e }); err != nil {
			return nil, err
		}
	}
	if query := piaq.withIPAddress; query != nil {
		if err := piaq.loadIPAddress(ctx, query, nodes, nil,
			func(n *PanierIPAddress, e *IPAddress) { n.Edges.IPAddress = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (piaq *PanierIPAddressQuery) loadPanier(ctx context.Context, query *PanierQuery, nodes []*PanierIPAddress, init func(*PanierIPAddress), assign func(*PanierIPAddress, *Panier)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierIPAddress)
	for i := range nodes {
		if nodes[i].panier_panier_ip_address == nil {
			continue
		}
		fk := *nodes[i].panier_panier_ip_address
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
			return fmt.Errorf(`unexpected foreign-key "panier_panier_ip_address" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (piaq *PanierIPAddressQuery) loadIPAddress(ctx context.Context, query *IPAddressQuery, nodes []*PanierIPAddress, init func(*PanierIPAddress), assign func(*PanierIPAddress, *IPAddress)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PanierIPAddress)
	for i := range nodes {
		if nodes[i]._IPAddress_id == nil {
			continue
		}
		fk := *nodes[i]._IPAddress_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ipaddress.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "_IPAddress_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (piaq *PanierIPAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piaq.querySpec()
	_spec.Node.Columns = piaq.ctx.Fields
	if len(piaq.ctx.Fields) > 0 {
		_spec.Unique = piaq.ctx.Unique != nil && *piaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piaq.driver, _spec)
}

func (piaq *PanierIPAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(panieripaddress.Table, panieripaddress.Columns, sqlgraph.NewFieldSpec(panieripaddress.FieldID, field.TypeInt))
	_spec.From = piaq.sql
	if unique := piaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piaq.path != nil {
		_spec.Unique = true
	}
	if fields := piaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, panieripaddress.FieldID)
		for i := range fields {
			if fields[i] != panieripaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piaq *PanierIPAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piaq.driver.Dialect())
	t1 := builder.Table(panieripaddress.Table)
	columns := piaq.ctx.Fields
	if len(columns) == 0 {
		columns = panieripaddress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piaq.sql != nil {
		selector = piaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piaq.ctx.Unique != nil && *piaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range piaq.predicates {
		p(selector)
	}
	for _, p := range piaq.order {
		p(selector)
	}
	if offset := piaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PanierIPAddressGroupBy is the group-by builder for PanierIPAddress entities.
type PanierIPAddressGroupBy struct {
	selector
	build *PanierIPAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (piagb *PanierIPAddressGroupBy) Aggregate(fns ...AggregateFunc) *PanierIPAddressGroupBy {
	piagb.fns = append(piagb.fns, fns...)
	return piagb
}

// Scan applies the selector query and scans the result into the given value.
func (piagb *PanierIPAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, piagb.build.ctx, "GroupBy")
	if err := piagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierIPAddressQuery, *PanierIPAddressGroupBy](ctx, piagb.build, piagb, piagb.build.inters, v)
}

func (piagb *PanierIPAddressGroupBy) sqlScan(ctx context.Context, root *PanierIPAddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(piagb.fns))
	for _, fn := range piagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*piagb.flds)+len(piagb.fns))
		for _, f := range *piagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*piagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := piagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PanierIPAddressSelect is the builder for selecting fields of PanierIPAddress entities.
type PanierIPAddressSelect struct {
	*PanierIPAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pias *PanierIPAddressSelect) Aggregate(fns ...AggregateFunc) *PanierIPAddressSelect {
	pias.fns = append(pias.fns, fns...)
	return pias
}

// Scan applies the selector query and scans the result into the given value.
func (pias *PanierIPAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pias.ctx, "Select")
	if err := pias.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PanierIPAddressQuery, *PanierIPAddressSelect](ctx, pias.PanierIPAddressQuery, pias, pias.inters, v)
}

func (pias *PanierIPAddressSelect) sqlScan(ctx context.Context, root *PanierIPAddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pias.fns))
	for _, fn := range pias.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pias.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pias.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}