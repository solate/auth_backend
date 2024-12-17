// pkg/ent/schema/user.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {

	return []ent.Field{
		field.String("username").Comment("用户名").Unique().NotEmpty(),
		field.String("password").Comment("密码").Sensitive().NotEmpty(),
		field.String("nickname").Comment("昵称").Optional(),
		field.String("email").Comment("邮箱").Optional(),
		field.String("phone").Comment("手机号").Unique().Optional(),
		field.Int8("status").Comment("状态").Default(1),
		field.Time("created_at").Comment("创建时间").Default(time.Now),
		field.Time("updated_at").Comment("更新时间").Default(time.Now).UpdateDefault(time.Now),
	}
}

// pkg/ent/schema/user.go
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Through("user_roles", UserRole.Type),
		edge.To("groups", UserGroup.Type).Through("user_group_users", UserGroupUser.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("phone"),
	}
}
