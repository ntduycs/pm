package services

import (
	"context"

	"github.com/samber/lo"
	"go.uber.org/zap"

	"project-management/ent"
	"project-management/mappers"
	"project-management/models"
	"project-management/repositories"
	"project-management/validators"
)

type PaPcService struct {
	logger                       *zap.Logger
	paPcRepository               *repositories.PaPcRepository
	paPcTechnicalScoreRepository *repositories.PaPcTechnicalScoreRepository
	memberRepository             *repositories.MemberRepository
	paPcMapper                   *mappers.PaPcMapper
	paPcTechnicalScoreMapper     *mappers.PaPcTechnicalScoreMapper
	memberMapper                 *mappers.MemberMapper
	requestValidator             *validators.RequestValidator
}

func NewPaPcService(
	logger *zap.Logger,
	paPcRepository *repositories.PaPcRepository,
	paPcTechnicalScoreRepository *repositories.PaPcTechnicalScoreRepository,
	memberRepository *repositories.MemberRepository,
	paPcMapper *mappers.PaPcMapper,
	paPcTechnicalScoreMapper *mappers.PaPcTechnicalScoreMapper,
	memberMapper *mappers.MemberMapper,
	requestValidator *validators.RequestValidator,
) *PaPcService {
	return &PaPcService{
		logger:                       logger,
		paPcRepository:               paPcRepository,
		paPcTechnicalScoreRepository: paPcTechnicalScoreRepository,
		memberRepository:             memberRepository,
		paPcMapper:                   paPcMapper,
		paPcTechnicalScoreMapper:     paPcTechnicalScoreMapper,
		memberMapper:                 memberMapper,
		requestValidator:             requestValidator,
	}
}

func (s *PaPcService) ListPaPcResults(ctx context.Context, req *models.ListPaPcResultsRequest) (*models.ListPaPcResultsResponse, error) {
	if err := s.requestValidator.Validate(req); err != nil {
		s.logger.Error("ListPaPcResults", err.ToZapFields()...)
		return nil, err
	}

	paPcLst, paPcCount, err := s.paPcRepository.FindAll(ctx, req)

	if err != nil {
		s.logger.Error("ListPaPcResults", zap.Error(err))
		return nil, err
	}

	memberIDs := make([]int, 0)

	for _, entity := range paPcLst {
		memberIDs = append(memberIDs, entity.MemberID)
	}

	memberIDs = lo.Uniq(memberIDs)

	memberLst, err := s.memberRepository.FindAllByIDIn(ctx, memberIDs)

	if err != nil {
		s.logger.Error("ListPaPcResults", zap.Error(err))
		return nil, err
	}

	memberByID := lo.SliceToMap(memberLst, func(item *ent.Member) (int, *models.Member) {
		return item.ID, s.memberMapper.ToModel(item)
	})

	paPcResults := make([]*models.PaPc, 0)

	for _, entity := range paPcLst {
		paPcResult := s.paPcMapper.ToModel(entity)
		paPcResult.Member = memberByID[entity.MemberID]

		paPcResults = append(paPcResults, paPcResult)
	}

	return &models.ListPaPcResultsResponse{
		Items: paPcResults,
		PageResponse: models.PageResponse{
			Page:  req.Page,
			Size:  req.Size,
			Total: paPcCount,
			Pages: paPcCount / req.Size,
			Count: len(paPcLst),
		},
	}, nil
}

func (s *PaPcService) UpsertPaPcResult(ctx context.Context, req *models.UpsertPaPcResultRequest) (*models.EmptyResponse, error) {
	if err := s.requestValidator.Validate(req); err != nil {
		s.logger.Error("UpsertPaPcResult", err.ToZapFields()...)
		return nil, err
	}

	_, err := s.paPcRepository.Save(ctx, req)

	if err != nil {
		s.logger.Error("UpsertPaPcResult", zap.Error(err))
		return nil, err
	}

	return &models.EmptyResponse{
		Message: "Update PA/PC result successfully",
	}, nil
}
