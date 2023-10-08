package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PaPc holds the schema definition for the PaPc entity.
type PaPc struct {
	ent.Schema
}

// Fields of the PaPc.
func (PaPc) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("member_id"),
		field.Float32("technical_score"),
		field.Float32("productivity_score"),
		field.Float32("collaboration_score"),
		field.Float32("development_score"),
		field.String("period"),
		field.String("note").Default(""),
	}
}

// Edges of the PaPc.
func (PaPc) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("member", Member.Type, "pa_pc_results", "member_id"),
		hasMany("technical_score_details", PaPcTechnicalScore.Type),
	}
}

// Annotations of the PaPc.
func (PaPc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("pa_pc"),
	}
}
