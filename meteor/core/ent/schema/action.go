package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.String("mode"),
		field.String("args"),
		field.Bool("queued"),
		field.Bool("responded"),
		field.String("result"),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("targeting", Host.Type).
			Ref("actions"),
	}
}
