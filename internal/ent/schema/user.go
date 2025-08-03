package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			NotEmpty().
			MaxLen(255),
		field.String("password_hash").
			NotEmpty().
			MaxLen(255),
		field.String("username").
			Unique().
			MaxLen(50).
			Optional(),
		field.String("first_name").
			MaxLen(100).
			Optional(),
		field.String("last_name").
			MaxLen(100).
			Optional(),
		field.String("current_level").
			MaxLen(10).
			Comment("A1, A2, B1, B2, C1, C2").
			Optional(),
		field.String("target_level").
			MaxLen(10).
			Optional(),
		field.Int("daily_goal").
			Default(20).
			Comment("Questions to review per day"),
		field.String("timezone").
			MaxLen(50).
			Default("UTC"),
		field.Time("last_active_at").
			Optional(),
		field.Bool("is_active").
			Default(true),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("decks", Deck.Type),
		edge.To("question_reviews", QuestionReview.Type),
		edge.To("question_attempts", QuestionAttempt.Type),
		edge.To("pronunciation_attempts", PronunciationAttempt.Type),
		edge.To("user_stats", UserStats.Type),
		edge.To("learning_streaks", LearningStreak.Type),
	}
}
