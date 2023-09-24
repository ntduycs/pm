package mappers

import (
	"strings"

	"project-management/ent"
	"project-management/models"
)

type MemberMapper struct {
}

func NewMemberMapper() *MemberMapper {
	return &MemberMapper{}
}

func (m *MemberMapper) ToModel(member *ent.Member) *models.Member {
	model := &models.Member{
		ID:          member.ID,
		Name:        member.Name,
		Email:       member.Email,
		Level:       string(member.Level),
		Positions:   strings.Split(member.Positions, ","),
		KPI:         member.Kpi,
		Category:    string(member.Category),
		TotalEffort: member.TotalEffort,
		Status:      string(member.Status),
	}

	if member.StartDate != nil {
		startDate := member.StartDate.Format("2006-01-02")
		model.StartDate = &startDate
	}

	if member.EndDate != nil {
		endDate := member.EndDate.Format("2006-01-02")
		model.EndDate = &endDate
	}

	return model
}

func (m *MemberMapper) ToModels(members []*ent.Member) []*models.Member {
	modelLst := make([]*models.Member, len(members))

	for i, member := range members {
		modelLst[i] = m.ToModel(member)
	}

	return modelLst
}
