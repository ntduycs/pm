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

// GetMember ...
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

// ListMembers ...
// @Summary List members
// @Tags    Member
// @Accept  json
// @Produce json
// @Param   page      query    int    false "Page"     default(1)
// @Param   size      query    int    false "Size"     default(20)
// @Param   category  query    string false "Category" Enums(official,buffer)
// @Param   positions query    string false "Positions"
// @Param   status    query    string false "Status" Enums(active,inactive,pending)
// @Success 200       {object} models.ListMembersResponse
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

// UpsertMember ...
// @Summary Upsert member
// @Tags    Member
// @Accept  json
// @Produce json
// @Param   body body     models.UpsertMemberRequest true "Upsert Member body"
// @Success 200  {object} models.EmptyResponse
// @Router  /members [post]
func (c *MemberController) UpsertMember(ctx *fiber.Ctx) error {
	req := &models.UpsertMemberRequest{}

	if err := ctx.BodyParser(req); err != nil {
		c.logger.Error("UpsertMember", zap.Error(err))
		return err
	}

	resp, err := c.memberService.UpsertMember(ctx.Context(), req)

	if err != nil {
		c.logger.Error("UpsertMember", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}

// DeleteMember ...
// @Summary Delete member
// @Tags    Member
// @Accept  json
// @Produce json
// @Param   id  path     int true "Member ID"
// @Success 200 {object} models.EmptyResponse
// @Router  /members/{id} [delete]
func (c *MemberController) DeleteMember(ctx *fiber.Ctx) error {
	req := &models.IDRequest{}

	if err := ctx.ParamsParser(req); err != nil {
		c.logger.Error("DeleteMember", zap.Error(err))
		return err
	}

	resp, err := c.memberService.DeleteMember(ctx.Context(), req)

	if err != nil {
		c.logger.Error("DeleteMember", zap.Error(err))
		return err
	}

	return ctx.JSON(resp)
}
