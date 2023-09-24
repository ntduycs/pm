package models

type Member struct {
	ID          int      `json:"id"`
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

func (r *ListMembersRequest) Norm() *ListMembersRequest {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Size < 1 {
		r.Size = 20
	}

	return r
}

type ListMembersResponse struct {
	PageResponse
	Items []*Member `json:"items"`
}

type GetMemberResponse struct {
	Item *Member `json:"item"`
}
