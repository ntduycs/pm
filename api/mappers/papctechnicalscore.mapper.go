package mappers

import (
	"project-management/ent"
	"project-management/models"
)

type PaPcTechnicalScoreMapper struct {
}

func NewPaPcTechnicalScoreMapper() *PaPcTechnicalScoreMapper {
	return &PaPcTechnicalScoreMapper{}
}

func (m *PaPcTechnicalScoreMapper) ToModel(paPcTechnicalScore *ent.PaPcTechnicalScore) *models.PaPcTechnicalScore {
	model := &models.PaPcTechnicalScore{
		ID:           paPcTechnicalScore.ID,
		PaPcId:       paPcTechnicalScore.PaPcID,
		Type:         string(paPcTechnicalScore.Type),
		Skill:        paPcTechnicalScore.Skill,
		SelfScore:    paPcTechnicalScore.SelfScore,
		PeerScore:    paPcTechnicalScore.PeerScore,
		ManagerScore: paPcTechnicalScore.ManagerScore,
	}

	return model
}

func (m *PaPcTechnicalScoreMapper) ToModels(paPcTechnicalScores []*ent.PaPcTechnicalScore) []*models.PaPcTechnicalScore {
	modelLst := make([]*models.PaPcTechnicalScore, len(paPcTechnicalScores))

	for i, paPcTechnicalScore := range paPcTechnicalScores {
		modelLst[i] = m.ToModel(paPcTechnicalScore)
	}

	return modelLst
}
