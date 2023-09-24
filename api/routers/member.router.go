package routers

import (
	"github.com/gofiber/fiber/v2"

	"project-management/controllers"
)

type MemberRouter struct {
	controller *controllers.MemberController
}

func NewMemberRouter(controller *controllers.MemberController) *MemberRouter {
	return &MemberRouter{
		controller: controller,
	}
}

func (r *MemberRouter) Register(app *fiber.App) {
	router := app.Group(r.GetPrefix())

	router.Get("/:id", r.controller.GetMember)
	router.Get("/", r.controller.ListMembers)
	router.Post("/", r.controller.UpsertMember)
	router.Delete("/:id", r.controller.DeleteMember)
}

func (r *MemberRouter) GetPrefix() string {
	return "/members"
}
