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
		field.String("no").Default("").Comment("编号"),
		field.Int("role").Range(1, 3).Comment("角色，1：超级管理员 2：代理商 3：商户 --  废弃"),
		field.String("name").NotEmpty().Default("").Comment("真实姓名"),
		field.String("phone").NotEmpty().Default("").Comment("电话"),
		field.String("email").Default("").Comment("邮箱"),
		field.Int("gender").Default(0).Comment("性别，1：男 2：女"),
		field.String("pwd_hashed").NotEmpty().Comment("hash后的密码"),
		field.String("pwd_salt").NotEmpty().Comment("密码加盐"),
		field.String("token").Default("").Comment("登录后的token信息"),
		field.Int("disable_status").Default(0).Comment("禁用状态，0：正常 1：禁用"),
		field.String("company").Default("").Comment("所属企业"),
		field.Int("parent_disable_status").Default(0).Comment("上级禁用状态，0：正常 1：禁用"),
		field.String("icon").Default("").Comment("用户头像"),
		field.Int("status").Default(1).Comment("状态: 1:启用, 2:禁用"),
		field.Time("created_at").Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone"),
	}
}
