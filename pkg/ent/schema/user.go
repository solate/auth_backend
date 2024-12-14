// pkg/ent/schema/user.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username").Unique().NotEmpty(),
		field.String("password").Sensitive().NotEmpty(),
		field.String("nickname").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.Int8("status").Default(1),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("roles", Role.Type).Through("user_roles", UserRole.Type),
	}
}
