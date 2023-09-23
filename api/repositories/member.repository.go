package repositories

import (
	"context"
	"database/sql"

	"project-management/ent"
	"project-management/ent/member"
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
