package jira

import "go.uber.org/zap"

type ProjectService interface {
	List() ([]Project, error)
	Get(id string) (Project, error)
}

type projectService struct {
	logger *zap.Logger
}

func (s *projectService) List() ([]Project, error) {
	// TODO implement me
	panic("implement me")
}

func (s *projectService) Get(id string) (Project, error) {
	// TODO implement me
	panic("implement me")
}

func NewProjectService(logger *zap.Logger) ProjectService {
	return &projectService{logger}
}
