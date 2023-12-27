// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/question"
	"poll-app/ent/questionoption"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionOptionCreate is the builder for creating a QuestionOption entity.
type QuestionOptionCreate struct {
	config
	mutation *QuestionOptionMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (qoc *QuestionOptionCreate) SetText(s string) *QuestionOptionCreate {
	qoc.mutation.SetText(s)
	return qoc
}

// SetCreatedAt sets the "created_at" field.
func (qoc *QuestionOptionCreate) SetCreatedAt(t time.Time) *QuestionOptionCreate {
	qoc.mutation.SetCreatedAt(t)
	return qoc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qoc *QuestionOptionCreate) SetNillableCreatedAt(t *time.Time) *QuestionOptionCreate {
	if t != nil {
		qoc.SetCreatedAt(*t)
	}
	return qoc
}

// SetUpdatedAt sets the "updated_at" field.
func (qoc *QuestionOptionCreate) SetUpdatedAt(t time.Time) *QuestionOptionCreate {
	qoc.mutation.SetUpdatedAt(t)
	return qoc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qoc *QuestionOptionCreate) SetNillableUpdatedAt(t *time.Time) *QuestionOptionCreate {
	if t != nil {
		qoc.SetUpdatedAt(*t)
	}
	return qoc
}

// SetQuestionID sets the "question_id" field.
func (qoc *QuestionOptionCreate) SetQuestionID(i int) *QuestionOptionCreate {
	qoc.mutation.SetQuestionID(i)
	return qoc
}

// SetNextOptionInvID sets the "next_option_inv" edge to the QuestionOption entity by ID.
func (qoc *QuestionOptionCreate) SetNextOptionInvID(id int) *QuestionOptionCreate {
	qoc.mutation.SetNextOptionInvID(id)
	return qoc
}

// SetNillableNextOptionInvID sets the "next_option_inv" edge to the QuestionOption entity by ID if the given value is not nil.
func (qoc *QuestionOptionCreate) SetNillableNextOptionInvID(id *int) *QuestionOptionCreate {
	if id != nil {
		qoc = qoc.SetNextOptionInvID(*id)
	}
	return qoc
}

// SetNextOptionInv sets the "next_option_inv" edge to the QuestionOption entity.
func (qoc *QuestionOptionCreate) SetNextOptionInv(q *QuestionOption) *QuestionOptionCreate {
	return qoc.SetNextOptionInvID(q.ID)
}

// SetNextOptionID sets the "next_option" edge to the QuestionOption entity by ID.
func (qoc *QuestionOptionCreate) SetNextOptionID(id int) *QuestionOptionCreate {
	qoc.mutation.SetNextOptionID(id)
	return qoc
}

// SetNillableNextOptionID sets the "next_option" edge to the QuestionOption entity by ID if the given value is not nil.
func (qoc *QuestionOptionCreate) SetNillableNextOptionID(id *int) *QuestionOptionCreate {
	if id != nil {
		qoc = qoc.SetNextOptionID(*id)
	}
	return qoc
}

// SetNextOption sets the "next_option" edge to the QuestionOption entity.
func (qoc *QuestionOptionCreate) SetNextOption(q *QuestionOption) *QuestionOptionCreate {
	return qoc.SetNextOptionID(q.ID)
}

// SetQuestion sets the "question" edge to the Question entity.
func (qoc *QuestionOptionCreate) SetQuestion(q *Question) *QuestionOptionCreate {
	return qoc.SetQuestionID(q.ID)
}

// Mutation returns the QuestionOptionMutation object of the builder.
func (qoc *QuestionOptionCreate) Mutation() *QuestionOptionMutation {
	return qoc.mutation
}

// Save creates the QuestionOption in the database.
func (qoc *QuestionOptionCreate) Save(ctx context.Context) (*QuestionOption, error) {
	qoc.defaults()
	return withHooks(ctx, qoc.sqlSave, qoc.mutation, qoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qoc *QuestionOptionCreate) SaveX(ctx context.Context) *QuestionOption {
	v, err := qoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qoc *QuestionOptionCreate) Exec(ctx context.Context) error {
	_, err := qoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qoc *QuestionOptionCreate) ExecX(ctx context.Context) {
	if err := qoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qoc *QuestionOptionCreate) defaults() {
	if _, ok := qoc.mutation.CreatedAt(); !ok {
		v := questionoption.DefaultCreatedAt()
		qoc.mutation.SetCreatedAt(v)
	}
	if _, ok := qoc.mutation.UpdatedAt(); !ok {
		v := questionoption.DefaultUpdatedAt()
		qoc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qoc *QuestionOptionCreate) check() error {
	if _, ok := qoc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "QuestionOption.text"`)}
	}
	if _, ok := qoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "QuestionOption.created_at"`)}
	}
	if _, ok := qoc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "QuestionOption.updated_at"`)}
	}
	if _, ok := qoc.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question_id", err: errors.New(`ent: missing required field "QuestionOption.question_id"`)}
	}
	if _, ok := qoc.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question", err: errors.New(`ent: missing required edge "QuestionOption.question"`)}
	}
	return nil
}

func (qoc *QuestionOptionCreate) sqlSave(ctx context.Context) (*QuestionOption, error) {
	if err := qoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	qoc.mutation.id = &_node.ID
	qoc.mutation.done = true
	return _node, nil
}

func (qoc *QuestionOptionCreate) createSpec() (*QuestionOption, *sqlgraph.CreateSpec) {
	var (
		_node = &QuestionOption{config: qoc.config}
		_spec = sqlgraph.NewCreateSpec(questionoption.Table, sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt))
	)
	if value, ok := qoc.mutation.Text(); ok {
		_spec.SetField(questionoption.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := qoc.mutation.CreatedAt(); ok {
		_spec.SetField(questionoption.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := qoc.mutation.UpdatedAt(); ok {
		_spec.SetField(questionoption.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := qoc.mutation.NextOptionInvIDs(); len(nodes) > 0 {
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
		_node.question_option_next_option = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qoc.mutation.NextOptionIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qoc.mutation.QuestionIDs(); len(nodes) > 0 {
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
		_node.QuestionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QuestionOptionCreateBulk is the builder for creating many QuestionOption entities in bulk.
type QuestionOptionCreateBulk struct {
	config
	err      error
	builders []*QuestionOptionCreate
}

// Save creates the QuestionOption entities in the database.
func (qocb *QuestionOptionCreateBulk) Save(ctx context.Context) ([]*QuestionOption, error) {
	if qocb.err != nil {
		return nil, qocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(qocb.builders))
	nodes := make([]*QuestionOption, len(qocb.builders))
	mutators := make([]Mutator, len(qocb.builders))
	for i := range qocb.builders {
		func(i int, root context.Context) {
			builder := qocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuestionOptionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, qocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qocb *QuestionOptionCreateBulk) SaveX(ctx context.Context) []*QuestionOption {
	v, err := qocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qocb *QuestionOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := qocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qocb *QuestionOptionCreateBulk) ExecX(ctx context.Context) {
	if err := qocb.Exec(ctx); err != nil {
		panic(err)
	}
}
