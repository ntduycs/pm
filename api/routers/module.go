package routers

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewServer(lifecycle fx.Lifecycle, logger *zap.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		// Override default configuration here if needed
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		EnablePrintRoutes: true,
		ColorScheme:       fiber.DefaultColors,
		RequestMethods:    fiber.DefaultMethods,
	})

	app.Use(fiberlogger.New(fiberlogger.Config{
		Format:   "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeZone: "Asia/Saigon",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status": "UP",
		})
	})

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
	NewServer,
)
