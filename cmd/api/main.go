package main

import (
	"fmt"
	"iago-effting/api-example/pkg/logs"
	"os"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http"
	"iago-effting/api-example/pkg/storage/database"
	"iago-effting/api-example/pkg/version"
)

func main() {
	os.Setenv("ENV", "dev")

	//var logger = logrus.New()
	//{
	//	logger.Out = os.Stdout
	//	logger.SetReportCaller(false)
	//
	//	logger.SetFormatter(&logrus.TextFormatter{
	//		ForceColors:      true,
	//		DisableColors:    false,
	//		DisableTimestamp: true,
	//		DisableSorting:   true,
	//		DisableQuote:     true,
	//	})
	//}
	logger := logs.NewLoggerService(logs.LogrusAdapter())

	logger.Debug("Env", os.Getenv("ENV"))

	configService := configs.NewConfigService(os.Getenv("ENV"), logger)
	configService.LoadEnvVars()

	_, dbError := database.StartConnection()
	if dbError != nil {
		logger.Error(dbError)
		os.Exit(-1)
	}

	logger.Debug("Env", configs.Env.Name)
	logger.Debug("Version", version.Version)
	logger.Debug("Date Release", version.Time)

	port := fmt.Sprintf(":%d", configs.Env.Server.Port)

	serverService := http.NewServerService(
		port,
		logger,
	)

	err := serverService.Run()
	if err != nil {
		logger.Error("Exit", err)
		os.Exit(-1)
	}
}
