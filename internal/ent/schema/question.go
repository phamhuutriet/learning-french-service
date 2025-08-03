package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("question_type").
			NotEmpty().
			MaxLen(50).
			Comment("english_to_french, french_to_english, pronunciation"),
		field.String("question_text").
			NotEmpty().
			Comment("The question displayed to user"),
		field.String("correct_answer").
			NotEmpty().
			Comment("Expected correct answer"),
		field.String("difficulty_level").
			MaxLen(10).
			Comment("Inherited from word or adjusted").
			Optional(),
		field.Time("generated_at").
			Default(time.Now),
		field.Bool("is_active").
			Default(true),
		field.JSON("options", []string{}).
			Comment("Multiple choice options if applicable").
			Optional(),
		field.String("hints").
			Comment("Helpful hints for the question").
			Optional(),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("word", Word.Type).
			Ref("questions").
			Unique().
			Required(),
		edge.To("question_reviews", QuestionReview.Type),
		edge.To("question_attempts", QuestionAttempt.Type),
		edge.To("pronunciation_attempts", PronunciationAttempt.Type),
	}
}
