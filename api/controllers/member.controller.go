package controllers

import (
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
