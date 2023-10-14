package jira

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"project-management/shared"
)

type permissionService struct {
	logger   *zap.Logger
	username string
	password string
}

func (s *permissionService) ListMyPermissions(projectKey string) (map[string]Permission, error) {
	url := shared.JiraServer + "/rest/api/2/mypermissions"

	var permissions Permissions

	permissions, err := shared.Get[Permissions](&shared.Config{
		BaseUrl: url,
		Method:  "GET",
		Headers: map[string]string{
			"Authorization": shared.ToBasicAuth(s.username, s.password),
		},
	}, permissions)

	if err != nil {
		s.logger.Error("Error while getting permissions", zap.Error(err))
	}

	for _, permission := range permissions.Permissions {
		fmt.Printf("%s - %s: %v\n", permission.Type, permission.Name, permission.HavePermission)
	}

	return permissions.Permissions, err
}

func NewPermissionService(logger *zap.Logger) PermissionService {
	return &permissionService{
		logger:   logger,
		username: viper.GetString("jira.username"),
		password: viper.GetString("jira.password"),
	}
}
