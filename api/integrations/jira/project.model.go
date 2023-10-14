package jira

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
