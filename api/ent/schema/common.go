package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// https://entgo.io/docs/schema-edges#o2o-two-types

func table(name string) schema.Annotation {
	return entsql.Annotation{
		Table: name,
	}
}

func belongToOne(from string, fromT any, ref string, column string) ent.Edge {
	return edge.From(from, fromT).
		Ref(ref).
		Unique().
		Field(column)
}

func mustBelongToOne(from string, fromT any, ref string, column string) ent.Edge {
	return edge.From(from, fromT).
		Ref(ref).
		Unique().
		Required().
		Field(column)
}

func belongToMany(from string, fromT any, ref string, column string) ent.Edge {
	return edge.From(from, fromT).
		Ref(ref).
		Field(column)
}

func belongToManyThrough(from string, fromT any, ref string, through string, throughT any) ent.Edge {
	return edge.From(from, fromT).
		Ref(ref).
		Through(through, throughT)
}

func hasMany(to string, toT any) ent.Edge {
	return edge.To(to, toT)
}

// hasManyThrough relationship owner side, used on tables are referenced to (not contain FK)
func hasManyThrough(to string, toT any, through string, throughT any) ent.Edge {
	return edge.To(to, toT).
		Through(through, throughT)
}

// hasOne relationship owner side, used on tables are referenced to (not contain FK)
func hasOne(to string, toT any) ent.Edge {
	return edge.To(to, toT).
		Unique()
}

func mustHaveOne(to string, toT any) ent.Edge {
	return edge.To(to, toT).
		Unique().
		Required()
}
