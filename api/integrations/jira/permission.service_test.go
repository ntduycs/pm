package jira

import (
	"testing"
)

var (
	permissionSvc = &permissionService{logger: logger, username: username, password: password}
)

func TestPermissionService_ListMyPermissions(t *testing.T) {
	// Test case 1: Successful permission list retrieval
	projectKey := "HRM"
	expectedPermissions := map[string]Permission{
		"ADMINISTER": {Key: "ADMINISTER", Name: "Administer projects"},
		"USER":       {Key: "USER", Name: "Browse projects and more"},
	}
	permissions, err := permissionSvc.ListMyPermissions(projectKey)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(permissions) != len(expectedPermissions) {
		t.Errorf("Expected %d permissions, but got %d", len(expectedPermissions), len(permissions))
	}
}
