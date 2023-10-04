package routers

import (
	"github.com/gofiber/fiber/v2"

	"project-management/controllers"
)

type PaPcRouter struct {
	controller *controllers.PaPcController
}

func NewPaPcRouter(controller *controllers.PaPcController) *PaPcRouter {
	return &PaPcRouter{
		controller: controller,
	}
}

func (r *PaPcRouter) Register(app *fiber.App) {
	router := app.Group(r.GetPrefix())

	router.Get("", r.controller.ListPaPcResults)
	router.Post("", r.controller.UpsertPaPcResult)
}

func (r *PaPcRouter) GetPrefix() string {
	return "/pa-pc-results"
}
