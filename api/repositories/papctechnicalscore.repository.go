package repositories

import (
	"context"
	"database/sql"

	"project-management/ent"
	"project-management/ent/papctechnicalscore"
)

type PaPcTechnicalScoreRepository struct {
	ent *ent.Client
	sql *sql.DB
}

func NewPaPcTechnicalScoreRepository(props RepositoryProps) *PaPcTechnicalScoreRepository {
	return &PaPcTechnicalScoreRepository{
		ent: props.Ent,
		sql: props.Sql,
	}
}

func (r *PaPcTechnicalScoreRepository) FindAllByPaPcID(ctx context.Context, paPcID int, txClient ...*ent.Client) ([]*ent.PaPcTechnicalScore, error) {
	client := useClient(r.ent, txClient...)

	return client.PaPcTechnicalScore.Query().
		Where(papctechnicalscore.PaPcIDEQ(paPcID)).
		All(ctx)
}
