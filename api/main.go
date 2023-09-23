package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"project-management/controllers"
	"project-management/repositories"
	"project-management/routers"
	"project-management/services"
)

func main() {
	if err := os.Setenv("TZ", "Asia/Saigon"); err != nil {
		panic(err)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	settings, _ := json.MarshalIndent(viper.AllSettings(), "", " \t")
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println(string(settings))
	fmt.Println(strings.Repeat("=", 10))

	app := fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
			}
		}),
		fx.Provide(func() *zap.Logger {
			logger, _ := zap.NewDevelopment()
			return logger
		}),
		repositories.Module,
		services.Module,
		controllers.Module,
		routers.Module,
		fx.Invoke(func(*fiber.App) {}),
	)

	app.Run()
}
