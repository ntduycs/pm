package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project-management/models"
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
	req := &models.IDRequest{}

	if err := ctx.ParamsParser(req); err != nil {
		c.logger.Error("GetMember", zap.Error(err))
		return err
	}

	resp, err := c.memberService.GetMember(ctx.Context(), req)

	if err != nil {
		c.logger.Error("GetMember", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}

// ListMembers
// @Summary List members
// @Tags    Member
// @Accept  json
// @Produce json
// @Param   page query    int false "Page" default(1)
// @Param   size query    int false "Size" default(20)
// @Success 200  {object} models.ListMembersResponse
// @Router  /members [get]
func (c *MemberController) ListMembers(ctx *fiber.Ctx) error {
	req := &models.ListMembersRequest{}

	if err := ctx.QueryParser(req); err != nil {
		c.logger.Error("ListMembers", zap.Error(err))
		return err
	}

	resp, err := c.memberService.ListMembers(ctx.Context(), req.Norm())

	if err != nil {
		c.logger.Error("ListMembers", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}
