package jira

import "go.uber.org/fx"

var Module = fx.Provide(
	NewProjectService,
)
