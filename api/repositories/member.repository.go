package repositories

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"project-management/ent"
	"project-management/ent/member"
	"project-management/models"
)

type MemberRepository struct {
	ent *ent.Client
	sql *sql.DB
}

func NewMemberRepository(props RepositoryProps) *MemberRepository {
	return &MemberRepository{
		ent: props.Ent,
		sql: props.Sql,
	}
}

func (r *MemberRepository) FindById(ctx context.Context, id int, txClient ...*ent.Client) (*ent.Member, error) {
	client := useClient(r.ent, txClient...)

	return client.Member.Query().
		Where(member.ID(id)).
		Only(ctx)
}

func (r *MemberRepository) FindAll(ctx context.Context, req *models.ListMembersRequest, txClient ...*ent.Client) ([]*ent.Member, int, error) {
	client := useClient(r.ent, txClient...)

	base := client.Member.Query()

	memberCount, err := base.Clone().Count(ctx)

	if err != nil {
		return []*ent.Member{}, 0, err
	}

	selectQuery := base.Clone().
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Order(ent.Asc(member.FieldName))

	memberLst, err := selectQuery.All(ctx)

	if err != nil {
		return []*ent.Member{}, 0, err
	}

	return memberLst, memberCount, nil
}

func (r *MemberRepository) Save(ctx context.Context, req *models.UpsertMemberRequest, txClient ...*ent.Client) (int, error) {
	client := useClient(r.ent, txClient...)

	cmd := client.Member.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetLevel(member.Level(req.Level)).
		SetPositions(strings.Join(req.Positions, ",")).
		SetKpi(req.KPI).
		SetCategory(member.Category(req.Category)).
		SetStatus(member.Status(req.Status)).
		SetTotalEffort(req.TotalEffort).
		OnConflictColumns(member.FieldEmail).
		UpdateNewValues()

	if req.StartDate != nil {
		startDate, _ := time.Parse("2006-01-02", *req.StartDate)
		cmd.SetStartDate(startDate)
	}

	if req.EndDate != nil {
		endDate, _ := time.Parse("2006-01-02", *req.EndDate)
		cmd.SetEndDate(endDate)
	}

	return cmd.ID(ctx)
}
