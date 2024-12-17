package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("角色名").NotEmpty(),
		field.String("code").Comment("角色编码").NotEmpty().Unique(),
		field.String("description").Comment("角色描述").Optional(),
		field.Int8("status").Comment("状态").Default(1),
		field.Time("created_at").Comment("创建时间").Default(time.Now),
		field.Time("updated_at").Comment("更新时间").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Role.// pkg/ent/schema/role.go
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type).Through("role_permissions", RolePermission.Type),
		edge.To("data_permissions", DataPermission.Type).Through("role_data_permissions", RoleDataPermission.Type),
		edge.From("users", User.Type).Ref("roles").Through("user_roles", UserRole.Type),
		edge.From("groups", UserGroup.Type).Ref("roles").Through("user_group_roles", UserGroupRole.Type),
	}
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}
