// Code generated by ent, DO NOT EDIT.

package startedpoll

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the startedpoll type in the database.
	Label = "started_poll"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPollID holds the string denoting the poll_id field in the database.
	FieldPollID = "poll_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePoll holds the string denoting the poll edge name in mutations.
	EdgePoll = "poll"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCompletedQuestions holds the string denoting the completed_questions edge name in mutations.
	EdgeCompletedQuestions = "completed_questions"
	// Table holds the table name of the startedpoll in the database.
	Table = "started_polls"
	// PollTable is the table that holds the poll relation/edge.
	PollTable = "started_polls"
	// PollInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollInverseTable = "polls"
	// PollColumn is the table column denoting the poll relation/edge.
	PollColumn = "poll_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "started_polls"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CompletedQuestionsTable is the table that holds the completed_questions relation/edge.
	CompletedQuestionsTable = "completed_questions"
	// CompletedQuestionsInverseTable is the table name for the CompletedQuestion entity.
	// It exists in this package in order to avoid circular dependency with the "completedquestion" package.
	CompletedQuestionsInverseTable = "completed_questions"
	// CompletedQuestionsColumn is the table column denoting the completed_questions relation/edge.
	CompletedQuestionsColumn = "started_poll_id"
)

// Columns holds all SQL columns for startedpoll fields.
var Columns = []string{
	FieldID,
	FieldPollID,
	FieldUserID,
	FieldCompleted,
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
	// DefaultCompleted holds the default value on creation for the "completed" field.
	DefaultCompleted bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the StartedPoll queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPollID orders the results by the poll_id field.
func ByPollID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPollID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCompleted orders the results by the completed field.
func ByCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompleted, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPollField orders the results by poll field.
func ByPollField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPollStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompletedQuestionsCount orders the results by completed_questions count.
func ByCompletedQuestionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompletedQuestionsStep(), opts...)
	}
}

// ByCompletedQuestions orders the results by completed_questions terms.
func ByCompletedQuestions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompletedQuestionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPollStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PollInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PollTable, PollColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newCompletedQuestionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompletedQuestionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CompletedQuestionsTable, CompletedQuestionsColumn),
	)
}
