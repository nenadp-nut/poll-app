// Code generated by ent, DO NOT EDIT.

package ent

import (
	"poll-app/ent/completedquestion"
	"poll-app/ent/poll"
	"poll-app/ent/question"
	"poll-app/ent/questionoption"
	"poll-app/ent/schema"
	"poll-app/ent/startedpoll"
	"poll-app/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	completedquestionFields := schema.CompletedQuestion{}.Fields()
	_ = completedquestionFields
	// completedquestionDescCreatedAt is the schema descriptor for created_at field.
	completedquestionDescCreatedAt := completedquestionFields[3].Descriptor()
	// completedquestion.DefaultCreatedAt holds the default value on creation for the created_at field.
	completedquestion.DefaultCreatedAt = completedquestionDescCreatedAt.Default.(func() time.Time)
	// completedquestionDescUpdatedAt is the schema descriptor for updated_at field.
	completedquestionDescUpdatedAt := completedquestionFields[4].Descriptor()
	// completedquestion.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	completedquestion.DefaultUpdatedAt = completedquestionDescUpdatedAt.Default.(func() time.Time)
	pollFields := schema.Poll{}.Fields()
	_ = pollFields
	// pollDescCreatedAt is the schema descriptor for created_at field.
	pollDescCreatedAt := pollFields[2].Descriptor()
	// poll.DefaultCreatedAt holds the default value on creation for the created_at field.
	poll.DefaultCreatedAt = pollDescCreatedAt.Default.(func() time.Time)
	// pollDescUpdatedAt is the schema descriptor for updated_at field.
	pollDescUpdatedAt := pollFields[3].Descriptor()
	// poll.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	poll.DefaultUpdatedAt = pollDescUpdatedAt.Default.(func() time.Time)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescHead is the schema descriptor for head field.
	questionDescHead := questionFields[2].Descriptor()
	// question.DefaultHead holds the default value on creation for the head field.
	question.DefaultHead = questionDescHead.Default.(bool)
	// questionDescRequired is the schema descriptor for required field.
	questionDescRequired := questionFields[3].Descriptor()
	// question.DefaultRequired holds the default value on creation for the required field.
	question.DefaultRequired = questionDescRequired.Default.(bool)
	// questionDescNumOfAnswers is the schema descriptor for num_of_answers field.
	questionDescNumOfAnswers := questionFields[4].Descriptor()
	// question.DefaultNumOfAnswers holds the default value on creation for the num_of_answers field.
	question.DefaultNumOfAnswers = questionDescNumOfAnswers.Default.(int)
	// questionDescCreatedAt is the schema descriptor for created_at field.
	questionDescCreatedAt := questionFields[5].Descriptor()
	// question.DefaultCreatedAt holds the default value on creation for the created_at field.
	question.DefaultCreatedAt = questionDescCreatedAt.Default.(func() time.Time)
	// questionDescUpdatedAt is the schema descriptor for updated_at field.
	questionDescUpdatedAt := questionFields[6].Descriptor()
	// question.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	question.DefaultUpdatedAt = questionDescUpdatedAt.Default.(func() time.Time)
	questionoptionFields := schema.QuestionOption{}.Fields()
	_ = questionoptionFields
	// questionoptionDescCreatedAt is the schema descriptor for created_at field.
	questionoptionDescCreatedAt := questionoptionFields[1].Descriptor()
	// questionoption.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionoption.DefaultCreatedAt = questionoptionDescCreatedAt.Default.(func() time.Time)
	// questionoptionDescUpdatedAt is the schema descriptor for updated_at field.
	questionoptionDescUpdatedAt := questionoptionFields[2].Descriptor()
	// questionoption.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	questionoption.DefaultUpdatedAt = questionoptionDescUpdatedAt.Default.(func() time.Time)
	startedpollFields := schema.StartedPoll{}.Fields()
	_ = startedpollFields
	// startedpollDescCompleted is the schema descriptor for completed field.
	startedpollDescCompleted := startedpollFields[2].Descriptor()
	// startedpoll.DefaultCompleted holds the default value on creation for the completed field.
	startedpoll.DefaultCompleted = startedpollDescCompleted.Default.(bool)
	// startedpollDescCreatedAt is the schema descriptor for created_at field.
	startedpollDescCreatedAt := startedpollFields[3].Descriptor()
	// startedpoll.DefaultCreatedAt holds the default value on creation for the created_at field.
	startedpoll.DefaultCreatedAt = startedpollDescCreatedAt.Default.(func() time.Time)
	// startedpollDescUpdatedAt is the schema descriptor for updated_at field.
	startedpollDescUpdatedAt := startedpollFields[4].Descriptor()
	// startedpoll.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	startedpoll.DefaultUpdatedAt = startedpollDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}
