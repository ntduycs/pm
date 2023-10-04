package mappers

import (
	"project-management/ent"
	"project-management/models"
)

type PaPcMapper struct {
}

func NewPaPcMapper() *PaPcMapper {
	return &PaPcMapper{}
}

func (m *PaPcMapper) ToModel(paPc *ent.PaPc) *models.PaPc {
	model := &models.PaPc{
		ID: paPc.ID,
		Member: &models.Member{
			ID: paPc.MemberID,
		},
		TechnicalScore:     paPc.TechnicalScore,
		ProductivityScore:  paPc.ProductivityScore,
		CollaborationScore: paPc.CollaborationScore,
		DevelopmentScore:   paPc.DevelopmentScore,
		Period:             paPc.Period,
	}

	return model
}

func (m *PaPcMapper) ToModels(paPcs []*ent.PaPc) []*models.PaPc {
	modelLst := make([]*models.PaPc, len(paPcs))

	for i, paPc := range paPcs {
		modelLst[i] = m.ToModel(paPc)
	}

	return modelLst
}
