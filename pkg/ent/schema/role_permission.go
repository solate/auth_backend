package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type RolePermission struct {
	ent.Schema
}

func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Int("role_id"),
		field.Int("permission_id"),
	}
}

func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Unique().
			Required().
			Field("role_id"),
		edge.To("permission", Permission.Type).
			Unique().
			Required().
			Field("permission_id"),
	}
}
