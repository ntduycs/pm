package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"project-management/controllers"
	"project-management/mappers"
	"project-management/repositories"
	"project-management/routers"
	"project-management/services"
	"project-management/validators"
)

// @title       Project Management API Document
// @version     1.0
// @description Project Management API Document

// @contact.name  Duy Nguyen
// @contact.email duy.nguyen-thanh@banvien.com.vn

// @host    localhost
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
		fx.Provide(func() *zap.Logger {
			config := zap.NewDevelopmentConfig()
			config.EncoderConfig.ConsoleSeparator = "  |  "
			config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(t.Format("2006-01-02 15:04:05"))
			}

			logger, _ := config.Build()
			return logger
		}),
		repositories.Module,
		mappers.Module,
		validators.Module,
		services.Module,
		controllers.Module,
		routers.Module,
		fx.Invoke(func(*fiber.App) {}),
	)

	app.Run()
}
