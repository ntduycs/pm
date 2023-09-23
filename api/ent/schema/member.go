package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"project-management/constants"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.Enum("level").Values(constants.MemberLevels...),
		field.String("positions").Comment("comma separated values of positions"),
		field.Float32("kpi").Min(0).Max(1).Comment("key performance indicator"),
		field.Enum("category").Values(constants.MemberCategories...),
		field.Float32("total_effort").Min(0).Max(100).Comment("total effort in percentage"),
		field.Time("start_date").Optional().Nillable().Comment("start date of member"),
		field.Time("end_date").Optional().Nillable().Comment("end date of member"),
		field.Enum("status").Values(constants.CommonStatuses...),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}

// Annotations of the Member.
func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("members"),
	}
}
