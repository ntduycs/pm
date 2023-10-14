package jira

type Permission struct {
	ID             string `json:"id"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	HavePermission bool   `json:"havePermission"`
}

type Permissions struct {
	Permissions map[string]Permission `json:"permissions"`
}
