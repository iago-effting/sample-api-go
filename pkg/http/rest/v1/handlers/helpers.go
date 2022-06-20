package handlers

import (
	"context"
	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/logs"
	"iago-effting/api-example/pkg/storage/database"
	"os"
)

func Setup() {
	os.Setenv("ENV", "test")
	logger := logs.NewLoggerService(logs.LogrusAdapter())

	configService := configs.NewConfigService(os.Getenv("ENV"), logger)
	configService.LoadEnvVars()

	database.StartConnection()
}

func ClearTable(table string) {
	ctx := context.Background()
	database.BunDb.NewTruncateTable().Table(table).Exec(ctx)
}
