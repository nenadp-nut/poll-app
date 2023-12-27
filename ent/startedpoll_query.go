// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"poll-app/ent/completedquestion"
	"poll-app/ent/poll"
	"poll-app/ent/predicate"
	"poll-app/ent/startedpoll"
	"poll-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StartedPollQuery is the builder for querying StartedPoll entities.
type StartedPollQuery struct {
	config
	ctx                    *QueryContext
	order                  []startedpoll.OrderOption
	inters                 []Interceptor
	predicates             []predicate.StartedPoll
	withPoll               *PollQuery
	withUser               *UserQuery
	withCompletedQuestions *CompletedQuestionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StartedPollQuery builder.
func (spq *StartedPollQuery) Where(ps ...predicate.StartedPoll) *StartedPollQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *StartedPollQuery) Limit(limit int) *StartedPollQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *StartedPollQuery) Offset(offset int) *StartedPollQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *StartedPollQuery) Unique(unique bool) *StartedPollQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *StartedPollQuery) Order(o ...startedpoll.OrderOption) *StartedPollQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryPoll chains the current query on the "poll" edge.
func (spq *StartedPollQuery) QueryPoll() *PollQuery {
	query := (&PollClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(startedpoll.Table, startedpoll.FieldID, selector),
			sqlgraph.To(poll.Table, poll.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, startedpoll.PollTable, startedpoll.PollColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (spq *StartedPollQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(startedpoll.Table, startedpoll.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, startedpoll.UserTable, startedpoll.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompletedQuestions chains the current query on the "completed_questions" edge.
func (spq *StartedPollQuery) QueryCompletedQuestions() *CompletedQuestionQuery {
	query := (&CompletedQuestionClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(startedpoll.Table, startedpoll.FieldID, selector),
			sqlgraph.To(completedquestion.Table, completedquestion.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, startedpoll.CompletedQuestionsTable, startedpoll.CompletedQuestionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StartedPoll entity from the query.
// Returns a *NotFoundError when no StartedPoll was found.
func (spq *StartedPollQuery) First(ctx context.Context) (*StartedPoll, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{startedpoll.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *StartedPollQuery) FirstX(ctx context.Context) *StartedPoll {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StartedPoll ID from the query.
// Returns a *NotFoundError when no StartedPoll ID was found.
func (spq *StartedPollQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{startedpoll.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *StartedPollQuery) FirstIDX(ctx context.Context) int {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StartedPoll entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StartedPoll entity is found.
// Returns a *NotFoundError when no StartedPoll entities are found.
func (spq *StartedPollQuery) Only(ctx context.Context) (*StartedPoll, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{startedpoll.Label}
	default:
		return nil, &NotSingularError{startedpoll.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *StartedPollQuery) OnlyX(ctx context.Context) *StartedPoll {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StartedPoll ID in the query.
// Returns a *NotSingularError when more than one StartedPoll ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *StartedPollQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{startedpoll.Label}
	default:
		err = &NotSingularError{startedpoll.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *StartedPollQuery) OnlyIDX(ctx context.Context) int {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StartedPolls.
func (spq *StartedPollQuery) All(ctx context.Context) ([]*StartedPoll, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*StartedPoll, *StartedPollQuery]()
	return withInterceptors[[]*StartedPoll](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *StartedPollQuery) AllX(ctx context.Context) []*StartedPoll {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StartedPoll IDs.
func (spq *StartedPollQuery) IDs(ctx context.Context) (ids []int, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(startedpoll.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *StartedPollQuery) IDsX(ctx context.Context) []int {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *StartedPollQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*StartedPollQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *StartedPollQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *StartedPollQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *StartedPollQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StartedPollQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *StartedPollQuery) Clone() *StartedPollQuery {
	if spq == nil {
		return nil
	}
	return &StartedPollQuery{
		config:                 spq.config,
		ctx:                    spq.ctx.Clone(),
		order:                  append([]startedpoll.OrderOption{}, spq.order...),
		inters:                 append([]Interceptor{}, spq.inters...),
		predicates:             append([]predicate.StartedPoll{}, spq.predicates...),
		withPoll:               spq.withPoll.Clone(),
		withUser:               spq.withUser.Clone(),
		withCompletedQuestions: spq.withCompletedQuestions.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithPoll tells the query-builder to eager-load the nodes that are connected to
// the "poll" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StartedPollQuery) WithPoll(opts ...func(*PollQuery)) *StartedPollQuery {
	query := (&PollClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withPoll = query
	return spq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StartedPollQuery) WithUser(opts ...func(*UserQuery)) *StartedPollQuery {
	query := (&UserClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withUser = query
	return spq
}

// WithCompletedQuestions tells the query-builder to eager-load the nodes that are connected to
// the "completed_questions" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StartedPollQuery) WithCompletedQuestions(opts ...func(*CompletedQuestionQuery)) *StartedPollQuery {
	query := (&CompletedQuestionClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withCompletedQuestions = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PollID int `json:"poll_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StartedPoll.Query().
//		GroupBy(startedpoll.FieldPollID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *StartedPollQuery) GroupBy(field string, fields ...string) *StartedPollGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StartedPollGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = startedpoll.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PollID int `json:"poll_id,omitempty"`
//	}
//
//	client.StartedPoll.Query().
//		Select(startedpoll.FieldPollID).
//		Scan(ctx, &v)
func (spq *StartedPollQuery) Select(fields ...string) *StartedPollSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &StartedPollSelect{StartedPollQuery: spq}
	sbuild.label = startedpoll.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StartedPollSelect configured with the given aggregations.
func (spq *StartedPollQuery) Aggregate(fns ...AggregateFunc) *StartedPollSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *StartedPollQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !startedpoll.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *StartedPollQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StartedPoll, error) {
	var (
		nodes       = []*StartedPoll{}
		_spec       = spq.querySpec()
		loadedTypes = [3]bool{
			spq.withPoll != nil,
			spq.withUser != nil,
			spq.withCompletedQuestions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StartedPoll).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StartedPoll{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withPoll; query != nil {
		if err := spq.loadPoll(ctx, query, nodes, nil,
			func(n *StartedPoll, e *Poll) { n.Edges.Poll = e }); err != nil {
			return nil, err
		}
	}
	if query := spq.withUser; query != nil {
		if err := spq.loadUser(ctx, query, nodes, nil,
			func(n *StartedPoll, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := spq.withCompletedQuestions; query != nil {
		if err := spq.loadCompletedQuestions(ctx, query, nodes,
			func(n *StartedPoll) { n.Edges.CompletedQuestions = []*CompletedQuestion{} },
			func(n *StartedPoll, e *CompletedQuestion) {
				n.Edges.CompletedQuestions = append(n.Edges.CompletedQuestions, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *StartedPollQuery) loadPoll(ctx context.Context, query *PollQuery, nodes []*StartedPoll, init func(*StartedPoll), assign func(*StartedPoll, *Poll)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*StartedPoll)
	for i := range nodes {
		fk := nodes[i].PollID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(poll.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "poll_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (spq *StartedPollQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*StartedPoll, init func(*StartedPoll), assign func(*StartedPoll, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*StartedPoll)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (spq *StartedPollQuery) loadCompletedQuestions(ctx context.Context, query *CompletedQuestionQuery, nodes []*StartedPoll, init func(*StartedPoll), assign func(*StartedPoll, *CompletedQuestion)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*StartedPoll)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(completedquestion.FieldStartedPollID)
	}
	query.Where(predicate.CompletedQuestion(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(startedpoll.CompletedQuestionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.StartedPollID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "started_poll_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (spq *StartedPollQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *StartedPollQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(startedpoll.Table, startedpoll.Columns, sqlgraph.NewFieldSpec(startedpoll.FieldID, field.TypeInt))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, startedpoll.FieldID)
		for i := range fields {
			if fields[i] != startedpoll.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if spq.withPoll != nil {
			_spec.Node.AddColumnOnce(startedpoll.FieldPollID)
		}
		if spq.withUser != nil {
			_spec.Node.AddColumnOnce(startedpoll.FieldUserID)
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *StartedPollQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(startedpoll.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = startedpoll.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StartedPollGroupBy is the group-by builder for StartedPoll entities.
type StartedPollGroupBy struct {
	selector
	build *StartedPollQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *StartedPollGroupBy) Aggregate(fns ...AggregateFunc) *StartedPollGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *StartedPollGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StartedPollQuery, *StartedPollGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *StartedPollGroupBy) sqlScan(ctx context.Context, root *StartedPollQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StartedPollSelect is the builder for selecting fields of StartedPoll entities.
type StartedPollSelect struct {
	*StartedPollQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *StartedPollSelect) Aggregate(fns ...AggregateFunc) *StartedPollSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *StartedPollSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StartedPollQuery, *StartedPollSelect](ctx, sps.StartedPollQuery, sps, sps.inters, v)
}

func (sps *StartedPollSelect) sqlScan(ctx context.Context, root *StartedPollQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}