// Code generated by ent, DO NOT EDIT.

package completedquestion

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the completedquestion type in the database.
	Label = "completed_question"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartedPollID holds the string denoting the started_poll_id field in the database.
	FieldStartedPollID = "started_poll_id"
	// FieldQuestionID holds the string denoting the question_id field in the database.
	FieldQuestionID = "question_id"
	// FieldAnswers holds the string denoting the answers field in the database.
	FieldAnswers = "answers"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeStartedPoll holds the string denoting the started_poll edge name in mutations.
	EdgeStartedPoll = "started_poll"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// Table holds the table name of the completedquestion in the database.
	Table = "completed_questions"
	// StartedPollTable is the table that holds the started_poll relation/edge.
	StartedPollTable = "completed_questions"
	// StartedPollInverseTable is the table name for the StartedPoll entity.
	// It exists in this package in order to avoid circular dependency with the "startedpoll" package.
	StartedPollInverseTable = "started_polls"
	// StartedPollColumn is the table column denoting the started_poll relation/edge.
	StartedPollColumn = "started_poll_id"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "completed_questions"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_id"
)

// Columns holds all SQL columns for completedquestion fields.
var Columns = []string{
	FieldID,
	FieldStartedPollID,
	FieldQuestionID,
	FieldAnswers,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the CompletedQuestion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStartedPollID orders the results by the started_poll_id field.
func ByStartedPollID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedPollID, opts...).ToFunc()
}

// ByQuestionID orders the results by the question_id field.
func ByQuestionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuestionID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStartedPollField orders the results by started_poll field.
func ByStartedPollField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStartedPollStep(), sql.OrderByField(field, opts...))
	}
}

// ByQuestionField orders the results by question field.
func ByQuestionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionStep(), sql.OrderByField(field, opts...))
	}
}
func newStartedPollStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StartedPollInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StartedPollTable, StartedPollColumn),
	)
}
func newQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
	)
}
