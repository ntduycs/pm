package integrations

import (
	"go.uber.org/fx"

	"project-management/integrations/jira"
)

var Module = fx.Module("integrations",
	jira.Module,
)
