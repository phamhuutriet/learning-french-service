package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserStats holds the schema definition for the UserStats entity.
type UserStats struct {
	ent.Schema
}

// Fields of the UserStats.
func (UserStats) Fields() []ent.Field {
	return []ent.Field{
		field.Time("stat_date").
			Default(time.Now),
		field.Int("questions_answered").
			Default(0),
		field.Int("questions_correct").
			Default(0),
		field.Int("words_learned").
			Default(0).
			Comment("Words with all 3 questions at 80%+ accuracy"),
		field.Int("total_study_time_minutes").
			Default(0),
		field.Int("streak_days").
			Default(0),
		field.Float("english_to_french_accuracy").
			Default(0.0),
		field.Float("french_to_english_accuracy").
			Default(0.0),
		field.Float("pronunciation_accuracy").
			Default(0.0),
		field.Int("average_response_time_ms").
			Default(0),
		field.Float("words_per_minute").
			Default(0.0),
		field.JSON("strengths", []string{}).
			Comment("AI-identified strengths").
			Optional(),
		field.JSON("weaknesses", []string{}).
			Comment("AI-identified areas for improvement").
			Optional(),
		field.String("recommendations").
			Comment("AI-generated study recommendations").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the UserStats.
func (UserStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_stats").
			Unique().
			Required(),
	}
}
