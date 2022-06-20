package main

import (
	"fmt"
	"os"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http"
	"iago-effting/api-example/pkg/storage/database"
	"iago-effting/api-example/pkg/version"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
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

	configService := configs.NewConfigService(os.Getenv("ENV"), logger)
	configService.LoadEnvVars()

	level.Debug(logger).Log("Env", configs.Env.Name)
	level.Debug(logger).Log("Version", version.Version)
	level.Debug(logger).Log("Date Release", version.Time)

	port := fmt.Sprintf(":%d", configs.Env.Server.Port)

	databaseService := database.NewDatabaseService(
		database.DatabaseOptions{
			DSN: configs.Env.Database.DSN,
		},
		logger,
	)

	err := databaseService.Connect()
	if err != nil {
		level.Error(logger).Log("Exit", err)
	}

	serverService := http.NewServerService(
		port,
		logger,
		databaseService.GetDb(),
	)

	err = serverService.Run()

	if err != nil {
		level.Error(logger).Log("Exit", err)
		os.Exit(-1)
	}
}
