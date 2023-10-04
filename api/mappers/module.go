package mappers

import "go.uber.org/fx"

var Module = fx.Provide(
	NewMemberMapper,
	NewPaPcMapper,
	NewPaPcTechnicalScoreMapper,
)
