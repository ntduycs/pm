package models

import (
	"github.com/samber/lo"

	"project-management/constants"
)

type Member struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	Level       constants.MemberLevel `json:"level"`
	Positions   []string              `json:"positions"`
	KPI         int                   `json:"kpi"`
	Category    string                `json:"category"`
	TotalEffort float32               `json:"total_effort"`
	StartDate   *string               `json:"start_date" format:"date"`
	EndDate     *string               `json:"end_date" format:"date"`
	Status      string                `json:"status"`
}

type ListMembersRequest struct {
	PageRequest
	Category  string   `query:"category"`
	Positions []string `query:"positions"`
	Status    string   `query:"status"`
}

func (r *ListMembersRequest) Norm() *ListMembersRequest {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Size < 1 {
		r.Size = 20
	}

	if !lo.Contains(constants.MemberCategories, r.Category) {
		r.Category = ""
	}

	if !lo.Contains(constants.CommonStatuses, r.Status) {
		r.Status = ""
	}

	if len(r.Positions) > 0 {
		r.Positions = lo.Filter(r.Positions, func(item string, _ int) bool {
			return lo.Contains(constants.MemberPositions, item)
		})
	}

	return r
}

type UpsertMemberRequest struct {
	ID          int                   `json:"id"`
	Email       string                `json:"email" validate:"required,email"`
	Name        string                `json:"name" validate:"required"`
	Level       constants.MemberLevel `json:"level" validate:"required"`
	Positions   []string              `json:"positions" validate:"required,min=1"`
	KPI         int                   `json:"kpi" validate:"required"`
	Category    string                `json:"category" validate:"required"`
	TotalEffort float32               `json:"total_effort" validate:"required"`
	StartDate   *string               `json:"start_date"`
	EndDate     *string               `json:"end_date"`
	Status      string                `json:"status" validate:"required"`
}

type ListMembersResponse struct {
	PageResponse
	Items []*Member `json:"items"`
}

type GetMemberResponse struct {
	Item *Member `json:"item"`
}
