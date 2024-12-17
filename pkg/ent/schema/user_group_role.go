package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserGroupRole struct {
	ent.Schema
}

func (UserGroupRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("group_id").Comment("用户组ID"),
		field.Int64("role_id").Comment("角色ID"),
		field.Time("created_at").Comment("创建时间").Default(time.Now),
		field.Time("updated_at").Comment("更新时间").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (UserGroupRole) Edges() []ent.Edge {
	return nil
}
