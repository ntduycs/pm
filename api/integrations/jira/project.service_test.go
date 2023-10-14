package jira

import (
	"testing"

	"github.com/samber/lo"
)

var (
	projectSvc = &projectService{logger: logger, username: username, password: password}
)

func TestProjectService_Get(t *testing.T) {
	// Test case 1: Successful project retrieval
	projectID := "HRM"
	expectedProject := Project{Name: "HCM", Key: projectID}
	project, err := projectSvc.Get(projectID)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if project.Name != expectedProject.Name {
		t.Errorf("Expected project %v, but got %v", expectedProject, project)
	}
}

func TestProjectService_List(t *testing.T) {
	// Test case 1: Successful project list retrieval
	expectedProjects := []Project{
		{Name: "HCM", Key: "HRM"},
		{Name: "HCM", Key: "HRM"},
		{Name: "HCM", Key: "HRM"},
	}
	projects, err := projectSvc.List(3)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(projects) != len(expectedProjects) {
		t.Errorf("Expected %d projects, but got %d", len(expectedProjects), len(projects))
	}

	if !lo.ContainsBy(projects, func(p Project) bool { return p.Name == "HCM" }) {
		t.Errorf("Expected projects to contain project with name %s", "HCM")
	}
}
