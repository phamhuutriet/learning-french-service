package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LearningStreak holds the schema definition for the LearningStreak entity.
type LearningStreak struct {
	ent.Schema
}

// Fields of the LearningStreak.
func (LearningStreak) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_date").
			Default(time.Now),
		field.Time("end_date").
			Optional(),
		field.Int("streak_length").
			Default(0),
		field.Bool("is_current").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the LearningStreak.
func (LearningStreak) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("learning_streaks").
			Unique().
			Required(),
	}
}
