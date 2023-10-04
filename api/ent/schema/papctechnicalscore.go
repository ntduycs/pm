package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PaPcTechnicalScore holds the schema definition for the PaPcTechnicalScore entity.
type PaPcTechnicalScore struct {
	ent.Schema
}

// Fields of the PaPcTechnicalScore.
func (PaPcTechnicalScore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("pa_pc_id"),
		field.Enum("type").Values("soft-skills", "hard-skills"),
		field.String("skill"),
		field.Float32("self_score"),
		field.Float32("peer_score").Optional().Nillable(),
		field.Float32("manager_score"),
	}
}

// Edges of the PaPcTechnicalScore.
func (PaPcTechnicalScore) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("pa_pc", PaPc.Type, "technical_score_details", "pa_pc_id"),
	}
}

// Annotations of the PaPcTechnicalScore.
func (PaPcTechnicalScore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("pa_pc_technical_score"),
	}
}
