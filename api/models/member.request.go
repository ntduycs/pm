package models

type Member struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	Level       string   `json:"level"`
	Positions   []string `json:"positions"`
	KPI         float32  `json:"kpi"`
	Category    string   `json:"category"`
	TotalEffort float32  `json:"total_effort"`
	StartDate   *string  `json:"start_date" format:"date"`
	EndDate     *string  `json:"end_date" format:"date"`
	Status      string   `json:"status"`
}

type ListMembersRequest struct {
	PageRequest
}

type ListMembersResponse struct {
	PageResponse
	Items []*Member `json:"items"`
}

type GetMemberResponse struct {
	Item *Member `json:"item"`
}
