package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entsql.WithComments(true),
		schema.Comment("用户"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty().Comment("用户名"),
		field.String("password").Sensitive().NotEmpty().Comment("密码"),
		field.String("nickname").Optional().Comment("昵称"),
		field.String("avatar").Optional().Comment("头像"),
		field.String("email").Optional().Comment("邮箱"),
		field.String("phone").Unique().Optional().Comment("手机号"),
		field.Int8("status").Default(1).Comment("状态: 1:启用, 2:禁用"),
		field.Time("created_at").Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Through("user_roles", UserRole.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("phone"),
	}
}
