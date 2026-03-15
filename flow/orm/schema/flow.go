package schema

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

func pkid() string {
	u, _ := uuid.NewV7()
	return u.String()
}

func char(size int) map[string]string {
	return map[string]string{
		dialect.MySQL:    fmt.Sprintf("char(%d)", size),
		dialect.SQLite:   fmt.Sprintf("char(%d)", size),
		dialect.Postgres: fmt.Sprintf("char(%d)", size),
	}
}
func varchar(size int) map[string]string {
	return map[string]string{
		dialect.MySQL:    fmt.Sprintf("varchar(%d)", size),
		dialect.SQLite:   fmt.Sprintf("varchar(%d)", size),
		dialect.Postgres: fmt.Sprintf("varchar(%d)", size),
	}
}

type Flow struct {
	ent.Schema
}

func (Flow) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().NotEmpty().DefaultFunc(pkid).SchemaType(char(36)),

		field.String("pk").Immutable().NotEmpty().SchemaType(varchar(64)).Comment("pk"),
		field.String("store_code").Immutable().NotEmpty().SchemaType(varchar(64)).Comment("编码"),
		field.String("store_name").Immutable().NotEmpty().SchemaType(varchar(64)).Comment("名称"),

		field.String("dt").Immutable().NotEmpty().SchemaType(varchar(64)),

		field.Int("in_total").Immutable().Default(0).Comment("总进入人数"),
	}
}

func (Flow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Flow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pk").Unique(),
		index.Fields("dt"),
		index.Fields("store_code"),
	}
}

func (Flow) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "g_flow_entry"},
	}
}
