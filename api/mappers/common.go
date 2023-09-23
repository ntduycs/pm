package mappers

import (
	"project-management/ent"
	"project-management/models"
)

type Mapper interface {
	ToModel(member *ent.Member) *models.Member
	ToModels(members []*ent.Member) []*models.Member
}
