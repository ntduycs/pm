package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project-management/services"
	"project-management/utils"
)

type EffortAllocController struct {
	logger             *zap.Logger
	effortAllocService *services.EffortAllocService
}

func NewEffortAllocController(
	logger *zap.Logger,
	effortAllocService *services.EffortAllocService,
) *EffortAllocController {
	return &EffortAllocController{
		logger:             logger,
		effortAllocService: effortAllocService,
	}
}

// UploadWeeklyReport ...
// @Summary Upload weekly report
// @Tags    Effort Allocation
// @Accept  multipart/form-data
// @Produce json
// @Param   file formData file true "File"
// @Success 200  {object} models.UploadEaWeeklyReportResponse
// @Router  /effort-allocation/weekly [post]
func (c *EffortAllocController) UploadWeeklyReport(ctx *fiber.Ctx) error {
	rows, err := utils.ReadExcelFile(ctx, "file")

	if err != nil {
		c.logger.Error("UploadWeeklyReport", zap.Error(err))
		return err
	}

	resp, err := c.effortAllocService.ParseWeeklyReport(ctx.Context(), rows)

	if err != nil {
		c.logger.Error("UploadWeeklyReport", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}
