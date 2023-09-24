package services

import (
	"context"

	"go.uber.org/zap"

	"project-management/mappers"
	"project-management/models"
	"project-management/repositories"
	"project-management/validators"
)

type MemberService struct {
	logger           *zap.Logger
	memberRepository *repositories.MemberRepository
	memberMapper     *mappers.MemberMapper
	requestValidator *validators.RequestValidator
}

func NewMemberService(
	logger *zap.Logger,
	memberRepository *repositories.MemberRepository,
	memberMapper *mappers.MemberMapper,
	requestValidator *validators.RequestValidator,
) *MemberService {
	return &MemberService{
		logger:           logger,
		memberRepository: memberRepository,
		memberMapper:     memberMapper,
		requestValidator: requestValidator,
	}
}

func (s *MemberService) GetMember(ctx context.Context, req *models.IDRequest) (*models.GetMemberResponse, error) {
	entity, err := s.memberRepository.FindById(ctx, req.ID)

	if err != nil {
		s.logger.Error("GetMember", zap.Error(err))
		return nil, err
	}

	return &models.GetMemberResponse{
		Item: s.memberMapper.ToModel(entity),
	}, nil
}

func (s *MemberService) ListMembers(ctx context.Context, req *models.ListMembersRequest) (*models.ListMembersResponse, error) {
	if err := s.requestValidator.Validate(req); err != nil {
		s.logger.Error("ListMembers", err.ToZapFields()...)
		return nil, err
	}

	entityLst, entityCount, err := s.memberRepository.FindAll(ctx, req)

	if err != nil {
		s.logger.Error("ListMembers", zap.Error(err))
		return nil, err
	}

	return &models.ListMembersResponse{
		Items: s.memberMapper.ToModels(entityLst),
		PageResponse: models.PageResponse{
			Page:  req.Page,
			Size:  req.Size,
			Total: entityCount,
			Pages: entityCount / req.Size,
		},
	}, nil
}

func (s *MemberService) CreateMember(ctx context.Context, req *models.UpsertMemberRequest) (*models.EmptyResponse, error) {
	if err := s.requestValidator.Validate(req); err != nil {
		s.logger.Error("CreateMember", err.ToZapFields()...)
		return nil, err
	}

	_, err := s.memberRepository.Save(ctx, req)

	if err != nil {
		s.logger.Error("CreateMember", zap.Error(err))
		return nil, err
	}

	return &models.EmptyResponse{
		Message: "Create member successfully",
	}, nil
}
