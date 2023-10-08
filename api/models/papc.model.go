package models

type PaPc struct {
	ID                 int     `json:"id"`
	Member             *Member `json:"member"`
	TechnicalScore     float32 `json:"technical_score"`
	ProductivityScore  float32 `json:"productivity_score"`
	CollaborationScore float32 `json:"collaboration_score"`
	DevelopmentScore   float32 `json:"development_score"`
	Period             string  `json:"period"`
	Note               string  `json:"note"`
}

type PaPcTechnicalScore struct {
	ID           int      `json:"id"`
	PaPcId       int      `json:"pa_pc_id"`
	Type         string   `json:"type"`
	Skill        string   `json:"skill"`
	SelfScore    float32  `json:"self_score"`
	PeerScore    *float32 `json:"peer_score"`
	ManagerScore float32  `json:"manager_score"`
}

type ListPaPcResultsRequest struct {
	PageRequest
	Period   string `query:"period"`
	MemberID int    `query:"member_id"`
}

func (r *ListPaPcResultsRequest) Norm() *ListPaPcResultsRequest {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Size < 1 {
		r.Size = 20
	}

	return r
}

type UpsertPaPcResultRequest struct {
	ID                 int     `json:"id"`
	MemberID           int     `json:"member_id" validate:"required"`
	TechnicalScore     float32 `json:"technical_score" validate:"required,gte=1"`
	ProductivityScore  float32 `json:"productivity_score" validate:"required,gte=1"`
	CollaborationScore float32 `json:"collaboration_score" validate:"required,gte=1"`
	DevelopmentScore   float32 `json:"development_score" validate:"required,gte=1"`
	Period             string  `json:"period" validate:"required"`
	Note               string  `json:"note"`
}

type ListPaPcResultsResponse struct {
	PageResponse
	Items []*PaPc `json:"items"`
}
