package repositories

import (
	"context"
	"database/sql"

	sql2 "entgo.io/ent/dialect/sql"
	"github.com/samber/lo"

	"project-management/ent"
	"project-management/ent/member"
	"project-management/ent/papc"
	"project-management/models"
)

type PaPcRepository struct {
	ent *ent.Client
	sql *sql.DB
}

func NewPaPcRepository(props RepositoryProps) *PaPcRepository {
	return &PaPcRepository{
		ent: props.Ent,
		sql: props.Sql,
	}
}

func (r *PaPcRepository) FindByID(ctx context.Context, id int, txClient ...*ent.Client) (*ent.PaPc, error) {
	client := useClient(r.ent, txClient...)

	return client.PaPc.Query().
		Where(papc.ID(id)).
		Only(ctx)
}

func (r *PaPcRepository) FindAll(ctx context.Context, req *models.ListPaPcResultsRequest, txClient ...*ent.Client) ([]*ent.PaPc, int, error) {
	client := useClient(r.ent, txClient...)

	base := client.PaPc.Query()

	if req.Period != "" {
		base = base.Where(papc.PeriodEQ(req.Period))
	}

	if req.MemberID > 0 {
		base = base.Where(papc.MemberIDEQ(req.MemberID))
	}

	total, err := base.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	paPcLst, err := base.
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Order(r.sort(req.Sort, req.Direction)...).
		All(ctx)

	if err != nil {
		return nil, 0, err
	}

	return paPcLst, total, nil
}

func (r *PaPcRepository) sort(sort string, direction string) []papc.OrderOption {
	orderTerm := lo.Ternary(direction == "descend", sql2.OrderDesc(), sql2.OrderAsc())

	switch sort {
	case "name":
		return []papc.OrderOption{
			papc.ByMemberField(member.FieldName, orderTerm),
		}
	default:
		return []papc.OrderOption{
			lo.Ternary(direction == "descend", ent.Desc(sort), ent.Asc(sort)),
		}
	}
}

func (r *PaPcRepository) Save(ctx context.Context, req *models.UpsertPaPcResultRequest, txClient ...*ent.Client) (int, error) {
	client := useClient(r.ent, txClient...)

	paPcID, err := client.PaPc.Create().
		SetMemberID(req.MemberID).
		SetPeriod(req.Period).
		SetTechnicalScore(req.TechnicalScore).
		SetProductivityScore(req.ProductivityScore).
		SetCollaborationScore(req.CollaborationScore).
		SetDevelopmentScore(req.DevelopmentScore).
		SetNote(req.Note).
		OnConflictColumns(papc.FieldMemberID, papc.FieldPeriod).
		UpdateNewValues().
		ID(ctx)

	if err != nil {
		return 0, err
	}

	return paPcID, nil
}
