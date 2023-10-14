package jira

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"project-management/shared"
)

type projectService struct {
	logger   *zap.Logger
	username string
	password string
}

func (s *projectService) List(recent int) ([]Project, error) {
	url := shared.JiraServer + "/rest/api/2/project"

	var projects []Project

	s.logger.Info("Getting projects", zap.String("recent", "3"))

	projects, err := shared.Get[[]Project](&shared.Config{
		BaseUrl: url,
		Method:  "GET",
		Headers: map[string]string{
			"Authorization": shared.ToBasicAuth(s.username, s.password),
		},
		Data: map[string]interface{}{
			"recent": recent,
		},
	}, projects)

	if err != nil {
		s.logger.Error("Error while getting projects", zap.Error(err))
	}

	return projects, err
}

func (s *projectService) Get(id string) (Project, error) {
	url := shared.JiraServer + "/rest/api/2/project/" + id

	var project Project

	s.logger.Info("Getting project", zap.String("id", id))

	project, err := shared.Get[Project](&shared.Config{
		BaseUrl: url,
		Method:  "GET",
		Headers: map[string]string{
			"Authorization": shared.ToBasicAuth(s.username, s.password),
		},
	}, project)

	if err != nil {
		s.logger.Error("Error while getting project", zap.String("id", id), zap.Error(err))
	}

	return project, err
}

func NewProjectService(logger *zap.Logger) ProjectService {
	return &projectService{
		logger:   logger,
		username: viper.GetString("jira.username"),
		password: viper.GetString("jira.password"),
	}
}
