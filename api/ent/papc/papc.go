// Code generated by ent, DO NOT EDIT.

package papc

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the papc type in the database.
	Label = "pa_pc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldTechnicalScore holds the string denoting the technical_score field in the database.
	FieldTechnicalScore = "technical_score"
	// FieldProductivityScore holds the string denoting the productivity_score field in the database.
	FieldProductivityScore = "productivity_score"
	// FieldCollaborationScore holds the string denoting the collaboration_score field in the database.
	FieldCollaborationScore = "collaboration_score"
	// FieldDevelopmentScore holds the string denoting the development_score field in the database.
	FieldDevelopmentScore = "development_score"
	// FieldPeriod holds the string denoting the period field in the database.
	FieldPeriod = "period"
	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "member"
	// EdgeTechnicalScoreDetails holds the string denoting the technical_score_details edge name in mutations.
	EdgeTechnicalScoreDetails = "technical_score_details"
	// Table holds the table name of the papc in the database.
	Table = "pa_pc"
	// MemberTable is the table that holds the member relation/edge.
	MemberTable = "pa_pc"
	// MemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MemberInverseTable = "member"
	// MemberColumn is the table column denoting the member relation/edge.
	MemberColumn = "member_id"
	// TechnicalScoreDetailsTable is the table that holds the technical_score_details relation/edge.
	TechnicalScoreDetailsTable = "pa_pc_technical_score"
	// TechnicalScoreDetailsInverseTable is the table name for the PaPcTechnicalScore entity.
	// It exists in this package in order to avoid circular dependency with the "papctechnicalscore" package.
	TechnicalScoreDetailsInverseTable = "pa_pc_technical_score"
	// TechnicalScoreDetailsColumn is the table column denoting the technical_score_details relation/edge.
	TechnicalScoreDetailsColumn = "pa_pc_id"
)

// Columns holds all SQL columns for papc fields.
var Columns = []string{
	FieldID,
	FieldMemberID,
	FieldTechnicalScore,
	FieldProductivityScore,
	FieldCollaborationScore,
	FieldDevelopmentScore,
	FieldPeriod,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the PaPc queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByTechnicalScore orders the results by the technical_score field.
func ByTechnicalScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTechnicalScore, opts...).ToFunc()
}

// ByProductivityScore orders the results by the productivity_score field.
func ByProductivityScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductivityScore, opts...).ToFunc()
}

// ByCollaborationScore orders the results by the collaboration_score field.
func ByCollaborationScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollaborationScore, opts...).ToFunc()
}

// ByDevelopmentScore orders the results by the development_score field.
func ByDevelopmentScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDevelopmentScore, opts...).ToFunc()
}

// ByPeriod orders the results by the period field.
func ByPeriod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPeriod, opts...).ToFunc()
}

// ByMemberField orders the results by member field.
func ByMemberField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberStep(), sql.OrderByField(field, opts...))
	}
}

// ByTechnicalScoreDetailsCount orders the results by technical_score_details count.
func ByTechnicalScoreDetailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTechnicalScoreDetailsStep(), opts...)
	}
}

// ByTechnicalScoreDetails orders the results by technical_score_details terms.
func ByTechnicalScoreDetails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTechnicalScoreDetailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMemberStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
	)
}
func newTechnicalScoreDetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TechnicalScoreDetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TechnicalScoreDetailsTable, TechnicalScoreDetailsColumn),
	)
}
