package database

import (
	"iago-effting/api-example/configs"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/uptrace/bun"
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
			DSN: configs.Env.Database.DSN,
		},
		logger,
	)

	err := databaseService.Connect()
	if err != nil {
		level.Error(logger).Log("Exit", err)
		os.Exit(-1)
	}

	return BunDb, nil
}
