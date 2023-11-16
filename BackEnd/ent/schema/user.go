package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

type Users struct {
    ent.Schema
}

func (Users) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
    }
}

func (Users) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("messages", Messages.Type).
            Ref("user").
            Unique(),
    }
}