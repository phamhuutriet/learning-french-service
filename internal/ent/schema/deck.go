package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Deck holds the schema definition for the Deck entity.
type Deck struct {
	ent.Schema
}

// Fields of the Deck.
func (Deck) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.String("description").
			Optional(),
		field.String("category").
			MaxLen(100).
			Comment("e.g., Travel, Business, Exam Prep").
			Optional(),
		field.String("target_level").
			MaxLen(10).
			Comment("A1, A2, B1, B2, C1, C2").
			Optional(),
		field.Bool("is_active").
			Default(true),
		field.Int("word_count").
			Default(0),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Deck.
func (Deck) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("decks").
			Unique().
			Required(),
		edge.To("words", Word.Type),
	}
}
