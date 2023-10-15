package routers

import (
	"github.com/gofiber/fiber/v2"

	"project-management/controllers"
)

type EffortAllocRouter struct {
	controller *controllers.EffortAllocController
}

func NewEffortAllocRouter(controller *controllers.EffortAllocController) *EffortAllocRouter {
	return &EffortAllocRouter{
		controller: controller,
	}
}

func (r *EffortAllocRouter) Register(app *fiber.App) {
	router := app.Group(r.GetPrefix())

	router.Post("/weekly", r.controller.UploadWeeklyReport)
}

func (r *EffortAllocRouter) GetPrefix() string {
	return "/effort-allocation"
}
