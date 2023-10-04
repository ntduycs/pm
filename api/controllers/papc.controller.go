package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project-management/models"
	"project-management/services"
)

type PaPcController struct {
	logger      *zap.Logger
	paPcService *services.PaPcService
}

func NewPaPcController(
	logger *zap.Logger,
	paPcService *services.PaPcService,
) *PaPcController {
	return &PaPcController{
		logger:      logger,
		paPcService: paPcService,
	}
}

// ListPaPcResults ...
// @Summary List PA/PC Results
// @Tags    PA/PC
// @Accept  json
// @Produce json
// @Param   page      query    int    false "Page" default(1)
// @Param   size      query    int    false "Size" default(20)
// @Param   member_id query    int    false "Member ID"
// @Param   period    query    string false "Period"
// @Success 200       {object} models.ListPaPcResultsResponse
// @Router  /pa-pc-results [get]
func (c *PaPcController) ListPaPcResults(ctx *fiber.Ctx) error {
	req := &models.ListPaPcResultsRequest{}

	if err := ctx.QueryParser(req); err != nil {
		c.logger.Error("ListPaPcResults", zap.Error(err))
		return err
	}

	resp, err := c.paPcService.ListPaPcResults(ctx.Context(), req.Norm())

	if err != nil {
		c.logger.Error("ListPaPcResults", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}

// UpsertPaPcResult ...
// @Summary Upsert PA/PC Result
// @Tags    PA/PC
// @Accept  json
// @Produce json
// @Param   body body     models.UpsertPaPcResultRequest true "Body"
// @Success 200  {object} models.EmptyResponse
// @Router  /pa-pc-results [post]
func (c *PaPcController) UpsertPaPcResult(ctx *fiber.Ctx) error {
	req := &models.UpsertPaPcResultRequest{}

	if err := ctx.BodyParser(req); err != nil {
		c.logger.Error("UpsertPaPcResult", zap.Error(err))
		return err
	}

	resp, err := c.paPcService.UpsertPaPcResult(ctx.Context(), req)

	if err != nil {
		c.logger.Error("UpsertPaPcResult", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}
