package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// pkg/ent/schema/permission.go
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("权限名称"),
		field.String("code").Unique().NotEmpty().Comment("权限编码"),
		field.Enum("type").Values("API", "MENU", "BUTTON", "DATA").Comment("权限类型"),
		field.String("resource_type").Optional().Comment("资源类型"),
		field.String("resource_path").Optional().Comment("资源路径"),
		field.Enum("action").Values("READ", "WRITE", "ALL").Optional().Comment("操作类型"),
		field.Int64("parent_id").Optional().Nillable().Comment("父权限ID"),
		field.String("description").Optional().Comment("描述"),
		field.Int8("status").Default(1).Comment("状态 1:启用 0:禁用"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Permission.Type).From("parent").Field("parent_id").Unique(),
		edge.From("roles", Role.Type).Ref("permissions").Through("role_permissions", RolePermission.Type),
	}
}
func (Permission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}
