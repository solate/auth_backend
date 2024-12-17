package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type RoleDataPermission struct {
	ent.Schema
}

func (RoleDataPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("role_id").Comment("角色ID"),
		field.Int64("data_permission_id").Comment("数据权限ID"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (RoleDataPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).Unique().Required().Field("role_id"),
		edge.To("data_permission", DataPermission.Type).Unique().Required().Field("data_permission_id"),
	}
}

func (RoleDataPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "data_permission_id").Unique(),
	}
}
