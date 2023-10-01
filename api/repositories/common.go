package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"project-management/ent"

	_ "github.com/lib/pq"
)

type RepositoryProps struct {
	fx.In
	Ent    *ent.Client
	Sql    *sql.DB
	Logger *zap.Logger
}

func database() string {
	user := viper.GetString("database.user")
	pass := viper.GetString("database.password")
	name := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")

	conn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		user, pass, name, host, port,
	)

	return conn
}

func useClient(defaultClient *ent.Client, txClient ...*ent.Client) *ent.Client {
	if len(txClient) > 0 {
		return txClient[0]
	}

	return defaultClient
}

func NewEntClient(lifecycle fx.Lifecycle, logger *zap.Logger) *ent.Client {
	options := func() []ent.Option {
		return []ent.Option{
			ent.Debug(),
			ent.Log(func(a ...any) {
				logger.Debug("ent", zap.Any("query", a))
			}),
		}
	}

	client, err := ent.Open("postgres", database(), options()...)

	if err != nil {
		logger.Fatal("failed opening connection to postgres", zap.Error(err))
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client
}

func NewSqlClient(lifecycle fx.Lifecycle, logger *zap.Logger) *sql.DB {
	db, err := sql.Open("postgres", database())

	if err != nil {
		logger.Fatal("failed opening connection to postgres", zap.Error(err))
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db
}
