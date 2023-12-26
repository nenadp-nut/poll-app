// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/predicate"
	"poll-app/ent/question"
	"poll-app/ent/questionoption"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionOptionUpdate is the builder for updating QuestionOption entities.
type QuestionOptionUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionOptionMutation
}

// Where appends a list predicates to the QuestionOptionUpdate builder.
func (qou *QuestionOptionUpdate) Where(ps ...predicate.QuestionOption) *QuestionOptionUpdate {
	qou.mutation.Where(ps...)
	return qou
}

// SetText sets the "text" field.
func (qou *QuestionOptionUpdate) SetText(s string) *QuestionOptionUpdate {
	qou.mutation.SetText(s)
	return qou
}

// SetNillableText sets the "text" field if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableText(s *string) *QuestionOptionUpdate {
	if s != nil {
		qou.SetText(*s)
	}
	return qou
}

// SetChosen sets the "chosen" field.
func (qou *QuestionOptionUpdate) SetChosen(b bool) *QuestionOptionUpdate {
	qou.mutation.SetChosen(b)
	return qou
}

// SetNillableChosen sets the "chosen" field if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableChosen(b *bool) *QuestionOptionUpdate {
	if b != nil {
		qou.SetChosen(*b)
	}
	return qou
}

// SetCreatedAt sets the "created_at" field.
func (qou *QuestionOptionUpdate) SetCreatedAt(t time.Time) *QuestionOptionUpdate {
	qou.mutation.SetCreatedAt(t)
	return qou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableCreatedAt(t *time.Time) *QuestionOptionUpdate {
	if t != nil {
		qou.SetCreatedAt(*t)
	}
	return qou
}

// SetUpdatedAt sets the "updated_at" field.
func (qou *QuestionOptionUpdate) SetUpdatedAt(t time.Time) *QuestionOptionUpdate {
	qou.mutation.SetUpdatedAt(t)
	return qou
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableUpdatedAt(t *time.Time) *QuestionOptionUpdate {
	if t != nil {
		qou.SetUpdatedAt(*t)
	}
	return qou
}

// SetQuestionID sets the "question_id" field.
func (qou *QuestionOptionUpdate) SetQuestionID(i int) *QuestionOptionUpdate {
	qou.mutation.SetQuestionID(i)
	return qou
}

// SetNillableQuestionID sets the "question_id" field if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableQuestionID(i *int) *QuestionOptionUpdate {
	if i != nil {
		qou.SetQuestionID(*i)
	}
	return qou
}

// SetNextOptionInvID sets the "next_option_inv" edge to the QuestionOption entity by ID.
func (qou *QuestionOptionUpdate) SetNextOptionInvID(id int) *QuestionOptionUpdate {
	qou.mutation.SetNextOptionInvID(id)
	return qou
}

// SetNillableNextOptionInvID sets the "next_option_inv" edge to the QuestionOption entity by ID if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableNextOptionInvID(id *int) *QuestionOptionUpdate {
	if id != nil {
		qou = qou.SetNextOptionInvID(*id)
	}
	return qou
}

// SetNextOptionInv sets the "next_option_inv" edge to the QuestionOption entity.
func (qou *QuestionOptionUpdate) SetNextOptionInv(q *QuestionOption) *QuestionOptionUpdate {
	return qou.SetNextOptionInvID(q.ID)
}

// SetNextOptionID sets the "next_option" edge to the QuestionOption entity by ID.
func (qou *QuestionOptionUpdate) SetNextOptionID(id int) *QuestionOptionUpdate {
	qou.mutation.SetNextOptionID(id)
	return qou
}

// SetNillableNextOptionID sets the "next_option" edge to the QuestionOption entity by ID if the given value is not nil.
func (qou *QuestionOptionUpdate) SetNillableNextOptionID(id *int) *QuestionOptionUpdate {
	if id != nil {
		qou = qou.SetNextOptionID(*id)
	}
	return qou
}

// SetNextOption sets the "next_option" edge to the QuestionOption entity.
func (qou *QuestionOptionUpdate) SetNextOption(q *QuestionOption) *QuestionOptionUpdate {
	return qou.SetNextOptionID(q.ID)
}

// SetQuestion sets the "question" edge to the Question entity.
func (qou *QuestionOptionUpdate) SetQuestion(q *Question) *QuestionOptionUpdate {
	return qou.SetQuestionID(q.ID)
}

// Mutation returns the QuestionOptionMutation object of the builder.
func (qou *QuestionOptionUpdate) Mutation() *QuestionOptionMutation {
	return qou.mutation
}

// ClearNextOptionInv clears the "next_option_inv" edge to the QuestionOption entity.
func (qou *QuestionOptionUpdate) ClearNextOptionInv() *QuestionOptionUpdate {
	qou.mutation.ClearNextOptionInv()
	return qou
}

// ClearNextOption clears the "next_option" edge to the QuestionOption entity.
func (qou *QuestionOptionUpdate) ClearNextOption() *QuestionOptionUpdate {
	qou.mutation.ClearNextOption()
	return qou
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qou *QuestionOptionUpdate) ClearQuestion() *QuestionOptionUpdate {
	qou.mutation.ClearQuestion()
	return qou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qou *QuestionOptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, qou.sqlSave, qou.mutation, qou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qou *QuestionOptionUpdate) SaveX(ctx context.Context) int {
	affected, err := qou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qou *QuestionOptionUpdate) Exec(ctx context.Context) error {
	_, err := qou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qou *QuestionOptionUpdate) ExecX(ctx context.Context) {
	if err := qou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qou *QuestionOptionUpdate) check() error {
	if _, ok := qou.mutation.QuestionID(); qou.mutation.QuestionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "QuestionOption.question"`)
	}
	return nil
}

func (qou *QuestionOptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := qou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(questionoption.Table, questionoption.Columns, sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt))
	if ps := qou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qou.mutation.Text(); ok {
		_spec.SetField(questionoption.FieldText, field.TypeString, value)
	}
	if value, ok := qou.mutation.Chosen(); ok {
		_spec.SetField(questionoption.FieldChosen, field.TypeBool, value)
	}
	if value, ok := qou.mutation.CreatedAt(); ok {
		_spec.SetField(questionoption.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := qou.mutation.UpdatedAt(); ok {
		_spec.SetField(questionoption.FieldUpdatedAt, field.TypeTime, value)
	}
	if qou.mutation.NextOptionInvCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   questionoption.NextOptionInvTable,
			Columns: []string{questionoption.NextOptionInvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qou.mutation.NextOptionInvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   questionoption.NextOptionInvTable,
			Columns: []string{questionoption.NextOptionInvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qou.mutation.NextOptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   questionoption.NextOptionTable,
			Columns: []string{questionoption.NextOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qou.mutation.NextOptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   questionoption.NextOptionTable,
			Columns: []string{questionoption.NextOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qou.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionoption.QuestionTable,
			Columns: []string{questionoption.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qou.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionoption.QuestionTable,
			Columns: []string{questionoption.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionoption.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qou.mutation.done = true
	return n, nil
}

// QuestionOptionUpdateOne is the builder for updating a single QuestionOption entity.
type QuestionOptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionOptionMutation
}

// SetText sets the "text" field.
func (qouo *QuestionOptionUpdateOne) SetText(s string) *QuestionOptionUpdateOne {
	qouo.mutation.SetText(s)
	return qouo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableText(s *string) *QuestionOptionUpdateOne {
	if s != nil {
		qouo.SetText(*s)
	}
	return qouo
}

// SetChosen sets the "chosen" field.
func (qouo *QuestionOptionUpdateOne) SetChosen(b bool) *QuestionOptionUpdateOne {
	qouo.mutation.SetChosen(b)
	return qouo
}

// SetNillableChosen sets the "chosen" field if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableChosen(b *bool) *QuestionOptionUpdateOne {
	if b != nil {
		qouo.SetChosen(*b)
	}
	return qouo
}

// SetCreatedAt sets the "created_at" field.
func (qouo *QuestionOptionUpdateOne) SetCreatedAt(t time.Time) *QuestionOptionUpdateOne {
	qouo.mutation.SetCreatedAt(t)
	return qouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableCreatedAt(t *time.Time) *QuestionOptionUpdateOne {
	if t != nil {
		qouo.SetCreatedAt(*t)
	}
	return qouo
}

// SetUpdatedAt sets the "updated_at" field.
func (qouo *QuestionOptionUpdateOne) SetUpdatedAt(t time.Time) *QuestionOptionUpdateOne {
	qouo.mutation.SetUpdatedAt(t)
	return qouo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableUpdatedAt(t *time.Time) *QuestionOptionUpdateOne {
	if t != nil {
		qouo.SetUpdatedAt(*t)
	}
	return qouo
}

// SetQuestionID sets the "question_id" field.
func (qouo *QuestionOptionUpdateOne) SetQuestionID(i int) *QuestionOptionUpdateOne {
	qouo.mutation.SetQuestionID(i)
	return qouo
}

// SetNillableQuestionID sets the "question_id" field if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableQuestionID(i *int) *QuestionOptionUpdateOne {
	if i != nil {
		qouo.SetQuestionID(*i)
	}
	return qouo
}

// SetNextOptionInvID sets the "next_option_inv" edge to the QuestionOption entity by ID.
func (qouo *QuestionOptionUpdateOne) SetNextOptionInvID(id int) *QuestionOptionUpdateOne {
	qouo.mutation.SetNextOptionInvID(id)
	return qouo
}

// SetNillableNextOptionInvID sets the "next_option_inv" edge to the QuestionOption entity by ID if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableNextOptionInvID(id *int) *QuestionOptionUpdateOne {
	if id != nil {
		qouo = qouo.SetNextOptionInvID(*id)
	}
	return qouo
}

// SetNextOptionInv sets the "next_option_inv" edge to the QuestionOption entity.
func (qouo *QuestionOptionUpdateOne) SetNextOptionInv(q *QuestionOption) *QuestionOptionUpdateOne {
	return qouo.SetNextOptionInvID(q.ID)
}

// SetNextOptionID sets the "next_option" edge to the QuestionOption entity by ID.
func (qouo *QuestionOptionUpdateOne) SetNextOptionID(id int) *QuestionOptionUpdateOne {
	qouo.mutation.SetNextOptionID(id)
	return qouo
}

// SetNillableNextOptionID sets the "next_option" edge to the QuestionOption entity by ID if the given value is not nil.
func (qouo *QuestionOptionUpdateOne) SetNillableNextOptionID(id *int) *QuestionOptionUpdateOne {
	if id != nil {
		qouo = qouo.SetNextOptionID(*id)
	}
	return qouo
}

// SetNextOption sets the "next_option" edge to the QuestionOption entity.
func (qouo *QuestionOptionUpdateOne) SetNextOption(q *QuestionOption) *QuestionOptionUpdateOne {
	return qouo.SetNextOptionID(q.ID)
}

// SetQuestion sets the "question" edge to the Question entity.
func (qouo *QuestionOptionUpdateOne) SetQuestion(q *Question) *QuestionOptionUpdateOne {
	return qouo.SetQuestionID(q.ID)
}

// Mutation returns the QuestionOptionMutation object of the builder.
func (qouo *QuestionOptionUpdateOne) Mutation() *QuestionOptionMutation {
	return qouo.mutation
}

// ClearNextOptionInv clears the "next_option_inv" edge to the QuestionOption entity.
func (qouo *QuestionOptionUpdateOne) ClearNextOptionInv() *QuestionOptionUpdateOne {
	qouo.mutation.ClearNextOptionInv()
	return qouo
}

// ClearNextOption clears the "next_option" edge to the QuestionOption entity.
func (qouo *QuestionOptionUpdateOne) ClearNextOption() *QuestionOptionUpdateOne {
	qouo.mutation.ClearNextOption()
	return qouo
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qouo *QuestionOptionUpdateOne) ClearQuestion() *QuestionOptionUpdateOne {
	qouo.mutation.ClearQuestion()
	return qouo
}

// Where appends a list predicates to the QuestionOptionUpdate builder.
func (qouo *QuestionOptionUpdateOne) Where(ps ...predicate.QuestionOption) *QuestionOptionUpdateOne {
	qouo.mutation.Where(ps...)
	return qouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (qouo *QuestionOptionUpdateOne) Select(field string, fields ...string) *QuestionOptionUpdateOne {
	qouo.fields = append([]string{field}, fields...)
	return qouo
}

// Save executes the query and returns the updated QuestionOption entity.
func (qouo *QuestionOptionUpdateOne) Save(ctx context.Context) (*QuestionOption, error) {
	return withHooks(ctx, qouo.sqlSave, qouo.mutation, qouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qouo *QuestionOptionUpdateOne) SaveX(ctx context.Context) *QuestionOption {
	node, err := qouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (qouo *QuestionOptionUpdateOne) Exec(ctx context.Context) error {
	_, err := qouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qouo *QuestionOptionUpdateOne) ExecX(ctx context.Context) {
	if err := qouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qouo *QuestionOptionUpdateOne) check() error {
	if _, ok := qouo.mutation.QuestionID(); qouo.mutation.QuestionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "QuestionOption.question"`)
	}
	return nil
}

func (qouo *QuestionOptionUpdateOne) sqlSave(ctx context.Context) (_node *QuestionOption, err error) {
	if err := qouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(questionoption.Table, questionoption.Columns, sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt))
	id, ok := qouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "QuestionOption.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := qouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questionoption.FieldID)
		for _, f := range fields {
			if !questionoption.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != questionoption.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := qouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qouo.mutation.Text(); ok {
		_spec.SetField(questionoption.FieldText, field.TypeString, value)
	}
	if value, ok := qouo.mutation.Chosen(); ok {
		_spec.SetField(questionoption.FieldChosen, field.TypeBool, value)
	}
	if value, ok := qouo.mutation.CreatedAt(); ok {
		_spec.SetField(questionoption.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := qouo.mutation.UpdatedAt(); ok {
		_spec.SetField(questionoption.FieldUpdatedAt, field.TypeTime, value)
	}
	if qouo.mutation.NextOptionInvCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   questionoption.NextOptionInvTable,
			Columns: []string{questionoption.NextOptionInvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qouo.mutation.NextOptionInvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   questionoption.NextOptionInvTable,
			Columns: []string{questionoption.NextOptionInvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qouo.mutation.NextOptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   questionoption.NextOptionTable,
			Columns: []string{questionoption.NextOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qouo.mutation.NextOptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   questionoption.NextOptionTable,
			Columns: []string{questionoption.NextOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qouo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionoption.QuestionTable,
			Columns: []string{questionoption.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qouo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionoption.QuestionTable,
			Columns: []string{questionoption.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &QuestionOption{config: qouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, qouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionoption.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	qouo.mutation.done = true
	return _node, nil
}