package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router interface {
	Register(app *fiber.App)
	GetPrefix() string
}

func asRouter(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}
