package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserGroup struct {
	ent.Schema
}

// pkg/ent/schema/user_group.go
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("用户组名称"),
		field.String("code").Unique().NotEmpty().Comment("用户组编码"),
		field.Int64("parent_id").Optional().Nillable().Comment("父级ID"),
		field.String("path").NotEmpty().Comment("用户组路径"),
		field.String("description").Optional().Comment("描述"),
		field.Int8("status").Default(1).Comment("状态 1:启用 0:禁用"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (UserGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Through("user_group_users", UserGroupUser.Type),
		edge.To("roles", Role.Type).Through("user_group_roles", UserGroupRole.Type),
		edge.To("children", UserGroup.Type).From("parent").Field("parent_id").Unique(),
	}
}

func (UserGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}
