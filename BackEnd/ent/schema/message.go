package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

type Messages struct {
    ent.Schema
}

func (Messages) Fields() []ent.Field {
    return []ent.Field{
        field.String("content"),
        field.Time("created_at").
            Default(time.Now).
            Immutable().
            Comment("Creation time of the message"),
        field.Int("USER_ID").
            Comment("ID of the user who sent the message"),
    }
}

func (Messages) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", Users.Type).
            Unique(),
    }
}