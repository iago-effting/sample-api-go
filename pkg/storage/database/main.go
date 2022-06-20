package database

import (
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"

	"iago-effting/api-example/configs"
)

func StartConnection() (*bun.DB, error) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	databaseService := NewDatabaseService(
		DatabaseOptions{
			DSN:      configs.Env.Database.DSN,
			User:     configs.Env.Database.User,
			Password: configs.Env.Database.Password,
			Host:     configs.Env.Database.Host,
			Port:     configs.Env.Database.Port | 5432,
			Name:     configs.Env.Database.Name,
		},
		logger,
	)

	err := databaseService.Connect()
	if err != nil {
		level.Error(logger).Log("Exit", err)
		os.Exit(-1)
	}

	if configs.Env.Debug.Database {
		BunDb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return BunDb, nil
}
