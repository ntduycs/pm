package repositories

import "go.uber.org/fx"

var Module = fx.Provide(
	NewSqlClient,
	NewEntClient,

	NewMemberRepository,
	NewPaPcRepository,
	NewPaPcTechnicalScoreRepository,
)

const (
	SortMemberName     = "name"
	SortMemberLevel    = "level"
	SortMemberKpi      = "kpi"
	SortMemberCategory = "category"
)
