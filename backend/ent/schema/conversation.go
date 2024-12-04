package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Conversation holds the schema definition for the Conversation entity.
type Conversation struct {
    ent.Schema
}

// Fields of the Conversation.
func (Conversation) Fields() []ent.Field {
    return []ent.Field{
        field.Text("messages").
            NotEmpty(),
        field.Time("created_at").
            Default(time.Now),
        field.String("user_id"),
    }
}

// Edges of the Conversation.
func (Conversation) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("prompt", Prompt.Type).
            Ref("conversations").
            Unique(),
    }
}
