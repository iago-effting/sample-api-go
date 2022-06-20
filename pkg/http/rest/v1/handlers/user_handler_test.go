package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/storage/database"
)

func init() {
	os.Setenv("ENV", "dev")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowError())
	}

	configService := configs.NewConfigService(os.Getenv("ENV"), logger)
	configService.LoadEnvVars()

	database.StartConnection()
}

func clearTables() {
	ctx := context.Background()
	database.BunDb.NewTruncateTable().Table("users").Exec(ctx)
}

func TestCreateUser(t *testing.T) {
	clearTables()

	server := gin.Default()
	server.POST("/v1/users", CreateUser)
	ts := httptest.NewServer(server)
	ts.Close()

	t.Run("User Created", func(t *testing.T) {
		params := CreateUserRequest{
			Email:          "test@test.com",
			Password:       "12345",
			RepeatPassword: "12345",
		}

		body, _ := json.Marshal(params)

		apitest.New().
			Handler(server).
			Post("/v1/users").
			JSON(string(body)).
			Expect(t).
			Assert(jsonpath.Equal(`$.data.email`, params.Email)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Params not valid", func(t *testing.T) {
		apitest.New().
			Handler(server).
			Post("/v1/users").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
}
