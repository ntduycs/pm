// Code generated by ent, DO NOT EDIT.

package member

import (
	"project-management/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldEmail, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldName, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldLevel, v))
}

// Positions applies equality check predicate on the "positions" field. It's identical to PositionsEQ.
func Positions(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldPositions, v))
}

// Kpi applies equality check predicate on the "kpi" field. It's identical to KpiEQ.
func Kpi(v int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldKpi, v))
}

// TotalEffort applies equality check predicate on the "total_effort" field. It's identical to TotalEffortEQ.
func TotalEffort(v float32) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldTotalEffort, v))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldEndDate, v))
}

// JiraName applies equality check predicate on the "jira_name" field. It's identical to JiraNameEQ.
func JiraName(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldJiraName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldEmail, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldName, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldLevel, v))
}

// PositionsEQ applies the EQ predicate on the "positions" field.
func PositionsEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldPositions, v))
}

// PositionsNEQ applies the NEQ predicate on the "positions" field.
func PositionsNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldPositions, v))
}

// PositionsIn applies the In predicate on the "positions" field.
func PositionsIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldPositions, vs...))
}

// PositionsNotIn applies the NotIn predicate on the "positions" field.
func PositionsNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldPositions, vs...))
}

// PositionsGT applies the GT predicate on the "positions" field.
func PositionsGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldPositions, v))
}

// PositionsGTE applies the GTE predicate on the "positions" field.
func PositionsGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldPositions, v))
}

// PositionsLT applies the LT predicate on the "positions" field.
func PositionsLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldPositions, v))
}

// PositionsLTE applies the LTE predicate on the "positions" field.
func PositionsLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldPositions, v))
}

// PositionsContains applies the Contains predicate on the "positions" field.
func PositionsContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldPositions, v))
}

// PositionsHasPrefix applies the HasPrefix predicate on the "positions" field.
func PositionsHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldPositions, v))
}

// PositionsHasSuffix applies the HasSuffix predicate on the "positions" field.
func PositionsHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldPositions, v))
}

// PositionsEqualFold applies the EqualFold predicate on the "positions" field.
func PositionsEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldPositions, v))
}

// PositionsContainsFold applies the ContainsFold predicate on the "positions" field.
func PositionsContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldPositions, v))
}

// KpiEQ applies the EQ predicate on the "kpi" field.
func KpiEQ(v int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldKpi, v))
}

// KpiNEQ applies the NEQ predicate on the "kpi" field.
func KpiNEQ(v int) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldKpi, v))
}

// KpiIn applies the In predicate on the "kpi" field.
func KpiIn(vs ...int) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldKpi, vs...))
}

// KpiNotIn applies the NotIn predicate on the "kpi" field.
func KpiNotIn(vs ...int) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldKpi, vs...))
}

// KpiGT applies the GT predicate on the "kpi" field.
func KpiGT(v int) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldKpi, v))
}

// KpiGTE applies the GTE predicate on the "kpi" field.
func KpiGTE(v int) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldKpi, v))
}

// KpiLT applies the LT predicate on the "kpi" field.
func KpiLT(v int) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldKpi, v))
}

// KpiLTE applies the LTE predicate on the "kpi" field.
func KpiLTE(v int) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldKpi, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldCategory, vs...))
}

// TotalEffortEQ applies the EQ predicate on the "total_effort" field.
func TotalEffortEQ(v float32) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldTotalEffort, v))
}

// TotalEffortNEQ applies the NEQ predicate on the "total_effort" field.
func TotalEffortNEQ(v float32) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldTotalEffort, v))
}

// TotalEffortIn applies the In predicate on the "total_effort" field.
func TotalEffortIn(vs ...float32) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldTotalEffort, vs...))
}

// TotalEffortNotIn applies the NotIn predicate on the "total_effort" field.
func TotalEffortNotIn(vs ...float32) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldTotalEffort, vs...))
}

// TotalEffortGT applies the GT predicate on the "total_effort" field.
func TotalEffortGT(v float32) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldTotalEffort, v))
}

// TotalEffortGTE applies the GTE predicate on the "total_effort" field.
func TotalEffortGTE(v float32) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldTotalEffort, v))
}

// TotalEffortLT applies the LT predicate on the "total_effort" field.
func TotalEffortLT(v float32) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldTotalEffort, v))
}

// TotalEffortLTE applies the LTE predicate on the "total_effort" field.
func TotalEffortLTE(v float32) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldTotalEffort, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldStartDate, v))
}

// StartDateIsNil applies the IsNil predicate on the "start_date" field.
func StartDateIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldStartDate))
}

// StartDateNotNil applies the NotNil predicate on the "start_date" field.
func StartDateNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldStartDate))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldEndDate, v))
}

// EndDateIsNil applies the IsNil predicate on the "end_date" field.
func EndDateIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldEndDate))
}

// EndDateNotNil applies the NotNil predicate on the "end_date" field.
func EndDateNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldEndDate))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldStatus, vs...))
}

// JiraNameEQ applies the EQ predicate on the "jira_name" field.
func JiraNameEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldJiraName, v))
}

// JiraNameNEQ applies the NEQ predicate on the "jira_name" field.
func JiraNameNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldJiraName, v))
}

// JiraNameIn applies the In predicate on the "jira_name" field.
func JiraNameIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldJiraName, vs...))
}

// JiraNameNotIn applies the NotIn predicate on the "jira_name" field.
func JiraNameNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldJiraName, vs...))
}

// JiraNameGT applies the GT predicate on the "jira_name" field.
func JiraNameGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldJiraName, v))
}

// JiraNameGTE applies the GTE predicate on the "jira_name" field.
func JiraNameGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldJiraName, v))
}

// JiraNameLT applies the LT predicate on the "jira_name" field.
func JiraNameLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldJiraName, v))
}

// JiraNameLTE applies the LTE predicate on the "jira_name" field.
func JiraNameLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldJiraName, v))
}

// JiraNameContains applies the Contains predicate on the "jira_name" field.
func JiraNameContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldJiraName, v))
}

// JiraNameHasPrefix applies the HasPrefix predicate on the "jira_name" field.
func JiraNameHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldJiraName, v))
}

// JiraNameHasSuffix applies the HasSuffix predicate on the "jira_name" field.
func JiraNameHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldJiraName, v))
}

// JiraNameEqualFold applies the EqualFold predicate on the "jira_name" field.
func JiraNameEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldJiraName, v))
}

// JiraNameContainsFold applies the ContainsFold predicate on the "jira_name" field.
func JiraNameContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldJiraName, v))
}

// HasPaPcResults applies the HasEdge predicate on the "pa_pc_results" edge.
func HasPaPcResults() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaPcResultsTable, PaPcResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaPcResultsWith applies the HasEdge predicate on the "pa_pc_results" edge with a given conditions (other predicates).
func HasPaPcResultsWith(preds ...predicate.PaPc) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newPaPcResultsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(sql.NotPredicates(p))
}
