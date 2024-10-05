// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"finance/ent/bank"
	"finance/ent/bank_transfer"
	"finance/ent/predicate"
	"finance/ent/transaction"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BankTransferQuery is the builder for querying Bank_Transfer entities.
type BankTransferQuery struct {
	config
	ctx              *QueryContext
	order            []bank_transfer.OrderOption
	inters           []Interceptor
	predicates       []predicate.Bank_Transfer
	withFromBank     *BankQuery
	withToBank       *BankQuery
	withTransactions *TransactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BankTransferQuery builder.
func (btq *BankTransferQuery) Where(ps ...predicate.Bank_Transfer) *BankTransferQuery {
	btq.predicates = append(btq.predicates, ps...)
	return btq
}

// Limit the number of records to be returned by this query.
func (btq *BankTransferQuery) Limit(limit int) *BankTransferQuery {
	btq.ctx.Limit = &limit
	return btq
}

// Offset to start from.
func (btq *BankTransferQuery) Offset(offset int) *BankTransferQuery {
	btq.ctx.Offset = &offset
	return btq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (btq *BankTransferQuery) Unique(unique bool) *BankTransferQuery {
	btq.ctx.Unique = &unique
	return btq
}

// Order specifies how the records should be ordered.
func (btq *BankTransferQuery) Order(o ...bank_transfer.OrderOption) *BankTransferQuery {
	btq.order = append(btq.order, o...)
	return btq
}

// QueryFromBank chains the current query on the "from_bank" edge.
func (btq *BankTransferQuery) QueryFromBank() *BankQuery {
	query := (&BankClient{config: btq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := btq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := btq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bank_transfer.Table, bank_transfer.FieldID, selector),
			sqlgraph.To(bank.Table, bank.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bank_transfer.FromBankTable, bank_transfer.FromBankColumn),
		)
		fromU = sqlgraph.SetNeighbors(btq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToBank chains the current query on the "to_bank" edge.
func (btq *BankTransferQuery) QueryToBank() *BankQuery {
	query := (&BankClient{config: btq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := btq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := btq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bank_transfer.Table, bank_transfer.FieldID, selector),
			sqlgraph.To(bank.Table, bank.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bank_transfer.ToBankTable, bank_transfer.ToBankColumn),
		)
		fromU = sqlgraph.SetNeighbors(btq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransactions chains the current query on the "transactions" edge.
func (btq *BankTransferQuery) QueryTransactions() *TransactionQuery {
	query := (&TransactionClient{config: btq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := btq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := btq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bank_transfer.Table, bank_transfer.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bank_transfer.TransactionsTable, bank_transfer.TransactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(btq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bank_Transfer entity from the query.
// Returns a *NotFoundError when no Bank_Transfer was found.
func (btq *BankTransferQuery) First(ctx context.Context) (*Bank_Transfer, error) {
	nodes, err := btq.Limit(1).All(setContextOp(ctx, btq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bank_transfer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (btq *BankTransferQuery) FirstX(ctx context.Context) *Bank_Transfer {
	node, err := btq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bank_Transfer ID from the query.
// Returns a *NotFoundError when no Bank_Transfer ID was found.
func (btq *BankTransferQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = btq.Limit(1).IDs(setContextOp(ctx, btq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bank_transfer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (btq *BankTransferQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := btq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bank_Transfer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bank_Transfer entity is found.
// Returns a *NotFoundError when no Bank_Transfer entities are found.
func (btq *BankTransferQuery) Only(ctx context.Context) (*Bank_Transfer, error) {
	nodes, err := btq.Limit(2).All(setContextOp(ctx, btq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bank_transfer.Label}
	default:
		return nil, &NotSingularError{bank_transfer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (btq *BankTransferQuery) OnlyX(ctx context.Context) *Bank_Transfer {
	node, err := btq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bank_Transfer ID in the query.
// Returns a *NotSingularError when more than one Bank_Transfer ID is found.
// Returns a *NotFoundError when no entities are found.
func (btq *BankTransferQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = btq.Limit(2).IDs(setContextOp(ctx, btq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bank_transfer.Label}
	default:
		err = &NotSingularError{bank_transfer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (btq *BankTransferQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := btq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bank_Transfers.
func (btq *BankTransferQuery) All(ctx context.Context) ([]*Bank_Transfer, error) {
	ctx = setContextOp(ctx, btq.ctx, ent.OpQueryAll)
	if err := btq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bank_Transfer, *BankTransferQuery]()
	return withInterceptors[[]*Bank_Transfer](ctx, btq, qr, btq.inters)
}

// AllX is like All, but panics if an error occurs.
func (btq *BankTransferQuery) AllX(ctx context.Context) []*Bank_Transfer {
	nodes, err := btq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bank_Transfer IDs.
func (btq *BankTransferQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if btq.ctx.Unique == nil && btq.path != nil {
		btq.Unique(true)
	}
	ctx = setContextOp(ctx, btq.ctx, ent.OpQueryIDs)
	if err = btq.Select(bank_transfer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (btq *BankTransferQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := btq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (btq *BankTransferQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, btq.ctx, ent.OpQueryCount)
	if err := btq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, btq, querierCount[*BankTransferQuery](), btq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (btq *BankTransferQuery) CountX(ctx context.Context) int {
	count, err := btq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (btq *BankTransferQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, btq.ctx, ent.OpQueryExist)
	switch _, err := btq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (btq *BankTransferQuery) ExistX(ctx context.Context) bool {
	exist, err := btq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BankTransferQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (btq *BankTransferQuery) Clone() *BankTransferQuery {
	if btq == nil {
		return nil
	}
	return &BankTransferQuery{
		config:           btq.config,
		ctx:              btq.ctx.Clone(),
		order:            append([]bank_transfer.OrderOption{}, btq.order...),
		inters:           append([]Interceptor{}, btq.inters...),
		predicates:       append([]predicate.Bank_Transfer{}, btq.predicates...),
		withFromBank:     btq.withFromBank.Clone(),
		withToBank:       btq.withToBank.Clone(),
		withTransactions: btq.withTransactions.Clone(),
		// clone intermediate query.
		sql:  btq.sql.Clone(),
		path: btq.path,
	}
}

// WithFromBank tells the query-builder to eager-load the nodes that are connected to
// the "from_bank" edge. The optional arguments are used to configure the query builder of the edge.
func (btq *BankTransferQuery) WithFromBank(opts ...func(*BankQuery)) *BankTransferQuery {
	query := (&BankClient{config: btq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	btq.withFromBank = query
	return btq
}

// WithToBank tells the query-builder to eager-load the nodes that are connected to
// the "to_bank" edge. The optional arguments are used to configure the query builder of the edge.
func (btq *BankTransferQuery) WithToBank(opts ...func(*BankQuery)) *BankTransferQuery {
	query := (&BankClient{config: btq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	btq.withToBank = query
	return btq
}

// WithTransactions tells the query-builder to eager-load the nodes that are connected to
// the "transactions" edge. The optional arguments are used to configure the query builder of the edge.
func (btq *BankTransferQuery) WithTransactions(opts ...func(*TransactionQuery)) *BankTransferQuery {
	query := (&TransactionClient{config: btq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	btq.withTransactions = query
	return btq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Amount float64 `json:"amount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BankTransfer.Query().
//		GroupBy(bank_transfer.FieldAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (btq *BankTransferQuery) GroupBy(field string, fields ...string) *BankTransferGroupBy {
	btq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BankTransferGroupBy{build: btq}
	grbuild.flds = &btq.ctx.Fields
	grbuild.label = bank_transfer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Amount float64 `json:"amount,omitempty"`
//	}
//
//	client.BankTransfer.Query().
//		Select(bank_transfer.FieldAmount).
//		Scan(ctx, &v)
func (btq *BankTransferQuery) Select(fields ...string) *BankTransferSelect {
	btq.ctx.Fields = append(btq.ctx.Fields, fields...)
	sbuild := &BankTransferSelect{BankTransferQuery: btq}
	sbuild.label = bank_transfer.Label
	sbuild.flds, sbuild.scan = &btq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BankTransferSelect configured with the given aggregations.
func (btq *BankTransferQuery) Aggregate(fns ...AggregateFunc) *BankTransferSelect {
	return btq.Select().Aggregate(fns...)
}

func (btq *BankTransferQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range btq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, btq); err != nil {
				return err
			}
		}
	}
	for _, f := range btq.ctx.Fields {
		if !bank_transfer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if btq.path != nil {
		prev, err := btq.path(ctx)
		if err != nil {
			return err
		}
		btq.sql = prev
	}
	return nil
}

func (btq *BankTransferQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bank_Transfer, error) {
	var (
		nodes       = []*Bank_Transfer{}
		_spec       = btq.querySpec()
		loadedTypes = [3]bool{
			btq.withFromBank != nil,
			btq.withToBank != nil,
			btq.withTransactions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bank_Transfer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bank_Transfer{config: btq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, btq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := btq.withFromBank; query != nil {
		if err := btq.loadFromBank(ctx, query, nodes, nil,
			func(n *Bank_Transfer, e *Bank) { n.Edges.FromBank = e }); err != nil {
			return nil, err
		}
	}
	if query := btq.withToBank; query != nil {
		if err := btq.loadToBank(ctx, query, nodes, nil,
			func(n *Bank_Transfer, e *Bank) { n.Edges.ToBank = e }); err != nil {
			return nil, err
		}
	}
	if query := btq.withTransactions; query != nil {
		if err := btq.loadTransactions(ctx, query, nodes,
			func(n *Bank_Transfer) { n.Edges.Transactions = []*Transaction{} },
			func(n *Bank_Transfer, e *Transaction) { n.Edges.Transactions = append(n.Edges.Transactions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (btq *BankTransferQuery) loadFromBank(ctx context.Context, query *BankQuery, nodes []*Bank_Transfer, init func(*Bank_Transfer), assign func(*Bank_Transfer, *Bank)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Bank_Transfer)
	for i := range nodes {
		fk := nodes[i].FromBankID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(bank.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "from_bank_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (btq *BankTransferQuery) loadToBank(ctx context.Context, query *BankQuery, nodes []*Bank_Transfer, init func(*Bank_Transfer), assign func(*Bank_Transfer, *Bank)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Bank_Transfer)
	for i := range nodes {
		fk := nodes[i].ToBankID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(bank.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "to_bank_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (btq *BankTransferQuery) loadTransactions(ctx context.Context, query *TransactionQuery, nodes []*Bank_Transfer, init func(*Bank_Transfer), assign func(*Bank_Transfer, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Bank_Transfer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(transaction.FieldTransferID)
	}
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bank_transfer.TransactionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TransferID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "transfer_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (btq *BankTransferQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := btq.querySpec()
	_spec.Node.Columns = btq.ctx.Fields
	if len(btq.ctx.Fields) > 0 {
		_spec.Unique = btq.ctx.Unique != nil && *btq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, btq.driver, _spec)
}

func (btq *BankTransferQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bank_transfer.Table, bank_transfer.Columns, sqlgraph.NewFieldSpec(bank_transfer.FieldID, field.TypeUUID))
	_spec.From = btq.sql
	if unique := btq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if btq.path != nil {
		_spec.Unique = true
	}
	if fields := btq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bank_transfer.FieldID)
		for i := range fields {
			if fields[i] != bank_transfer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if btq.withFromBank != nil {
			_spec.Node.AddColumnOnce(bank_transfer.FieldFromBankID)
		}
		if btq.withToBank != nil {
			_spec.Node.AddColumnOnce(bank_transfer.FieldToBankID)
		}
	}
	if ps := btq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := btq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := btq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := btq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (btq *BankTransferQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(btq.driver.Dialect())
	t1 := builder.Table(bank_transfer.Table)
	columns := btq.ctx.Fields
	if len(columns) == 0 {
		columns = bank_transfer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if btq.sql != nil {
		selector = btq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if btq.ctx.Unique != nil && *btq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range btq.predicates {
		p(selector)
	}
	for _, p := range btq.order {
		p(selector)
	}
	if offset := btq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := btq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BankTransferGroupBy is the group-by builder for Bank_Transfer entities.
type BankTransferGroupBy struct {
	selector
	build *BankTransferQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (btgb *BankTransferGroupBy) Aggregate(fns ...AggregateFunc) *BankTransferGroupBy {
	btgb.fns = append(btgb.fns, fns...)
	return btgb
}

// Scan applies the selector query and scans the result into the given value.
func (btgb *BankTransferGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, btgb.build.ctx, ent.OpQueryGroupBy)
	if err := btgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BankTransferQuery, *BankTransferGroupBy](ctx, btgb.build, btgb, btgb.build.inters, v)
}

func (btgb *BankTransferGroupBy) sqlScan(ctx context.Context, root *BankTransferQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(btgb.fns))
	for _, fn := range btgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*btgb.flds)+len(btgb.fns))
		for _, f := range *btgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*btgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := btgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BankTransferSelect is the builder for selecting fields of BankTransfer entities.
type BankTransferSelect struct {
	*BankTransferQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bts *BankTransferSelect) Aggregate(fns ...AggregateFunc) *BankTransferSelect {
	bts.fns = append(bts.fns, fns...)
	return bts
}

// Scan applies the selector query and scans the result into the given value.
func (bts *BankTransferSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bts.ctx, ent.OpQuerySelect)
	if err := bts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BankTransferQuery, *BankTransferSelect](ctx, bts.BankTransferQuery, bts, bts.inters, v)
}

func (bts *BankTransferSelect) sqlScan(ctx context.Context, root *BankTransferQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bts.fns))
	for _, fn := range bts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
