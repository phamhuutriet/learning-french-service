package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PronunciationAttempt holds the schema definition for the PronunciationAttempt entity.
type PronunciationAttempt struct {
	ent.Schema
}

// Fields of the PronunciationAttempt.
func (PronunciationAttempt) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_audio_url").
			NotEmpty().
			MaxLen(500).
			Comment("User recording"),
		field.String("reference_audio_url").
			MaxLen(500).
			Comment("Reference pronunciation").
			Optional(),
		field.Float("similarity_score").
			Comment("0.0 to 1.0 comparison with reference").
			Optional(),
		field.String("pronunciation_feedback").
			Comment("AI-generated pronunciation tips").
			Optional(),
		field.JSON("phonetic_accuracy", map[string]interface{}{}).
			Comment("Per-phoneme accuracy scores").
			Optional(),
		field.Time("attempt_date").
			Default(time.Now),
		field.Bool("was_successful").
			Comment("User satisfied with attempt").
			Optional(),
	}
}

// Edges of the PronunciationAttempt.
func (PronunciationAttempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("pronunciation_attempts").
			Unique().
			Required(),
		edge.From("question", Question.Type).
			Ref("pronunciation_attempts").
			Unique().
			Required(),
		edge.From("word", Word.Type).
			Ref("pronunciation_attempts").
			Unique().
			Required(),
	}
}
