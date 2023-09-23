package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type RouterProps struct {
	App    *fiber.App
	Prefix string
}

type Router interface {
	Register(props *RouterProps)
	GetPrefix() string
}

func asRouter(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}
