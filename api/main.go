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
	"project-management/integrations"
	"project-management/mappers"
	"project-management/repositories"
	"project-management/routers"
	"project-management/services"
	"project-management/shared"
	"project-management/validators"
)

// @title       Project Management API Document
// @version     1.0
// @description Project Management API Document

// @contact.name  Duy Nguyen
// @contact.email ntduy.cs@gmail.com

// @host    localhost:4000
// @schemes http https
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

	settings, _ := json.MarshalIndent(viper.AllSettings(), "", "   ")
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println(string(settings))
	fmt.Println(strings.Repeat("=", 10))

	app := fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
			}
		}),
		shared.Module,
		repositories.Module,
		mappers.Module,
		validators.Module,
		services.Module,
		integrations.Module,
		controllers.Module,
		routers.Module,
		fx.Invoke(func(*fiber.App) {}),
	)

	app.Run()
}
