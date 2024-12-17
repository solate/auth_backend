package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserGroupUser struct {
	ent.Schema
}

func (UserGroupUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户ID"),
		field.Int64("group_id").Comment("用户组ID"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (UserGroupUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("group", UserGroup.Type).Unique().Required().Field("group_id"),
	}
}

func (UserGroupUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "group_id").Unique(),
	}
}
