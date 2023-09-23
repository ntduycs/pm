package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project-management/services"
)

type MemberController struct {
	logger        *zap.Logger
	memberService *services.MemberService
}

func NewMemberController(
	logger *zap.Logger,
	memberService *services.MemberService,
) *MemberController {
	return &MemberController{
		logger:        logger,
		memberService: memberService,
	}
}

// GetMember
// @Summary Get member
// @Tags    Member
// @Accept  json
// @Produce json
// @Param   id  path     int true "Member ID"
// @Success 200 {object} models.GetMemberResponse
// @Router  /members/{id} [get]
func (c *MemberController) GetMember(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		c.logger.Error("GetMember", zap.Error(err))
		return err
	}

	resp, err := c.memberService.GetMember(ctx.Context(), int32(id))

	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}
