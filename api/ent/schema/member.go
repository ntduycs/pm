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
		field.String("email").Unique().NotEmpty().Comment("email address of member"),
		field.String("name"),
		field.Int("level").Min(0).Max(12).Comment("level of member"),
		field.String("positions").Comment("comma-separated sorted values of positions"),
		field.Int("kpi").Min(0).Max(200).Comment("key performance indicator"),
		field.Enum("category").Values(constants.MemberCategories...),
		field.Float32("total_effort").Min(0).Max(100).Comment("total effort in percentage"),
		field.Time("start_date").Optional().Nillable().Comment("start date of member"),
		field.Time("end_date").Optional().Nillable().Comment("end date of member"),
		field.Enum("status").Values(constants.CommonStatuses...),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		hasMany("pa_pc_results", PaPc.Type),
	}
}

// Annotations of the Member.
func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("member"),
	}
}
