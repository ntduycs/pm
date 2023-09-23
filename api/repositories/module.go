package repositories

import "go.uber.org/fx"

var Module = fx.Provide(
	NewSqlClient,
	NewEntClient,

	NewMemberRepository,
)
