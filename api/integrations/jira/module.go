package jira

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	username = ""
	password = ""
)

var (
	logger = zap.NewNop()
)

var Module = fx.Provide(
	NewProjectService,
	NewPermissionService,
)

type (
	ProjectService interface {
		List(recent int) ([]Project, error)
		Get(id string) (Project, error)
	}

	PermissionService interface {
		ListMyPermissions(projectKey string) (map[string]Permission, error)
	}
)
