package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// QuestionAttempt holds the schema definition for the QuestionAttempt entity.
type QuestionAttempt struct {
	ent.Schema
}

// Fields of the QuestionAttempt.
func (QuestionAttempt) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_answer").
			Optional(),
		field.String("ai_grade").
			MaxLen(20).
			Comment("correct, close, incorrect").
			Optional(),
		field.String("ai_feedback").
			Comment("AI-generated feedback").
			Optional(),
		field.Float("similarity_score").
			Comment("0.0 to 1.0 for semantic similarity").
			Optional(),
		field.Int("quality_rating").
			Comment("0=Again, 1=Hard, 2=Good, 3=Easy").
			Optional(),
		field.Int("response_time_ms").
			Comment("Time taken to answer in milliseconds").
			Optional(),
		field.Bool("was_correct").
			Default(false),
		field.Bool("needs_repeat").
			Default(false),
		field.Time("attempt_date").
			Default(time.Now),
	}
}

// Edges of the QuestionAttempt.
func (QuestionAttempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("question_attempts").
			Unique().
			Required(),
		edge.From("question", Question.Type).
			Ref("question_attempts").
			Unique().
			Required(),
		edge.From("question_review", QuestionReview.Type).
			Ref("question_attempts").
			Unique().
			Required(),
	}
}
