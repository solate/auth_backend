package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// func (UserRole) Annotations() []schema.Annotation {
// 	return []schema.Annotation{
// 		field.ID("user_id", "role_id"),
// 	}
// }

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Int("user_id"),
		field.Int("role_id"),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("role", Role.Type).
			Unique().
			Required().
			Field("role_id"),
	}
}
