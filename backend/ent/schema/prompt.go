package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Prompt holds the schema definition for the Prompt entity.
type Prompt struct {
    ent.Schema
}

// Fields of the Prompt.
func (Prompt) Fields() []ent.Field {
    return []ent.Field{
        field.String("title").
            NotEmpty(),
        field.Text("content").
            NotEmpty(),
        field.String("category"),
        field.Int("votes").
            Default(0),
        field.Time("created_at").
            Default(time.Now),
        field.String("created_by"),
    }
}

// Edges of the Prompt.
func (Prompt) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("conversations", Conversation.Type),
    }
}
