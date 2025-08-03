package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// QuestionReview holds the schema definition for the QuestionReview entity.
type QuestionReview struct {
	ent.Schema
}

// Fields of the QuestionReview.
func (QuestionReview) Fields() []ent.Field {
	return []ent.Field{
		field.Float("ease_factor").
			Default(2.5).
			Comment("SM-2 ease factor (1.3 minimum)"),
		field.Int("interval_days").
			Default(1).
			Comment("Days until next review"),
		field.Int("repetition_count").
			Default(0).
			Comment("Number of successful reviews"),
		field.Time("next_review_date").
			Default(time.Now),
		field.Time("last_reviewed_at").
			Optional(),
		field.Bool("is_due").
			Default(true),
		field.Int("total_attempts").
			Default(0),
		field.Int("correct_attempts").
			Default(0),
		field.Float("accuracy_rate").
			Default(0.0).
			Comment("correct_attempts / total_attempts"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the QuestionReview.
func (QuestionReview) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("question_reviews").
			Unique().
			Required(),
		edge.From("question", Question.Type).
			Ref("question_reviews").
			Unique().
			Required(),
		edge.To("question_attempts", QuestionAttempt.Type),
	}
}
