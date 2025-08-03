package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("french_word").
			NotEmpty().
			MaxLen(255).
			Comment("User provided"),
		field.String("english_translation").
			Comment("AI generated"),
		field.String("part_of_speech").
			MaxLen(50).
			Comment("AI generated: noun, verb, adjective, etc.").
			Optional(),
		field.String("gender").
			MaxLen(10).
			Comment("AI generated: masculine, feminine, neutral").
			Optional(),
		field.String("french_example").
			Comment("AI generated example sentence").
			Optional(),
		field.String("english_example_translation").
			Comment("AI generated translation of example").
			Optional(),
		field.String("phonetic_transcription").
			MaxLen(255).
			Comment("IPA notation from trusted API").
			Optional(),
		field.String("pronunciation_audio_url").
			MaxLen(500).
			Comment("URL to pronunciation audio file").
			Optional(),
		field.String("difficulty_level").
			MaxLen(10).
			Comment("A1, A2, B1, B2, C1, C2").
			Optional(),
		field.JSON("tags", []string{}).
			Comment("Array of user tags for filtering").
			Optional(),
		field.String("usage_context").
			Comment("When/how to use this word").
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

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deck", Deck.Type).
			Ref("words").
			Unique().
			Required(),
		edge.To("questions", Question.Type),
		edge.To("pronunciation_attempts", PronunciationAttempt.Type),
	}
}
