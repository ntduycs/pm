package services

import (
	"go.uber.org/zap"

	"project-management/repositories"
)

type MemberService struct {
	logger           *zap.Logger
	memberRepository *repositories.MemberRepository
}

func NewMemberService(
	logger *zap.Logger,
	memberRepository *repositories.MemberRepository,
) *MemberService {
	return &MemberService{
		logger:           logger,
		memberRepository: memberRepository,
	}
}
