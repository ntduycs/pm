package repositories

import (
	"context"
	"database/sql"

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

func (r *MemberRepository) FindById(ctx context.Context, id int32, txClient ...*ent.Client) (*ent.Member, error) {
	client := useClient(r.ent, txClient...)

	return client.Member.Query().
		Where(member.ID(id)).
		Only(ctx)
}

func (r *MemberRepository) FindAll(ctx context.Context, req *models.ListMembersRequest, txClient ...*ent.Client) ([]*ent.Member, int32, error) {
	client := useClient(r.ent, txClient...)

	base := client.Member.Query()

	memberCount, err := base.Clone().Count(ctx)

	if err != nil {
		return []*ent.Member{}, 0, err
	}

	selectQuery := base.Clone().
		Offset(int((req.Page - 1) * req.Size)).
		Limit(int(req.Size)).
		Order(ent.Asc(member.FieldName))

	memberLst, err := selectQuery.All(ctx)

	if err != nil {
		return []*ent.Member{}, 0, err
	}

	return memberLst, int32(memberCount), nil
}
