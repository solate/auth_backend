package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DataPermission struct {
	ent.Schema
}

func (DataPermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("数据权限名称"),
		field.String("resource_type").NotEmpty().Comment("资源类型"),
		field.String("condition_type").NotEmpty().Comment("条件类型: SQL/EXPR"),
		field.String("condition").NotEmpty().Comment("条件表达式"),
		field.String("description").Optional().Comment("描述"),
		field.Int8("status").Default(1).Comment("状态 1:启用 0:禁用"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// pkg/ent/schema/data_permission.go
func (DataPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("data_permissions").Through("role_data_permissions", RoleDataPermission.Type),
	}
}
