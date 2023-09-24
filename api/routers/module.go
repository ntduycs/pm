package routers

import (
	"context"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	fibercors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"project-management/ent"
	pmerror "project-management/errors"
	"project-management/models"

	"github.com/gofiber/swagger"

	_ "project-management/docs"
)

type ServerProps struct {
	fx.In
	fx.Lifecycle
	Logger  *zap.Logger
	Routers []Router `group:"routers"`
}

func NewServer(props ServerProps) *fiber.App {
	lifecycle := props.Lifecycle
	logger := props.Logger
	routers := props.Routers

	app := fiber.New(fiber.Config{
		// Override default configuration here if needed
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		RequestMethods: []string{
			fiber.MethodGet,
			fiber.MethodHead,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions,
		},
		EnablePrintRoutes: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var fe *fiber.Error
			var ve *pmerror.ValidatorError
			if errors.As(err, &ve) {
				return ctx.Status(fiber.StatusBadRequest).JSON(ve.ToErrorResponse())
			} else if errors.As(err, &fe) {
				code = fe.Code
			} else if ent.IsNotFound(err) {
				return ctx.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
					Message: strings.TrimPrefix(err.Error(), "ent: "),
				})
			} else if ent.IsConstraintError(err) {
				return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
					Message: strings.TrimPrefix(err.Error(), "ent: "),
				})
			} else if ent.IsValidationError(err) {
				return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
					Message: strings.TrimPrefix(err.Error(), "ent: "),
				})
			}

			return ctx.Status(code).JSON(models.ErrorResponse{
				Message: err.Error(),
			})
		},
	})

	app.Use(fibercors.New(fibercors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
		}, ","),
		MaxAge: 3600,
	}))
	app.Use(fiberlogger.New(fiberlogger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Saigon",
	}))
	app.Use(fiberrecover.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status": "UP",
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	for _, router := range routers {
		router.Register(app)
	}

	swagger.New(swagger.ConfigDefault)

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Server starting")

			go func() {
				address := fmt.Sprintf(":%d", viper.GetInt("server.port"))
				if err := app.Listen(address); err != nil {
					logger.Fatal("Server error", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Server stopped")
			return app.ShutdownWithContext(ctx)
		},
	})

	return app
}

var Module = fx.Provide(
	asRouter(NewMemberRouter),
	NewServer,
)
