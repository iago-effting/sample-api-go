package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/storage/database"
)

func init() {
	os.Setenv("ENV", "test")
	var logger = logrus.New()
	{
		logger.Out = os.Stdout
		logger.SetReportCaller(false)

		logger.SetFormatter(&logrus.TextFormatter{
			DisableColors:    true,
			DisableTimestamp: true,
			DisableSorting:   true,
			DisableQuote:     true,
		})
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
	server.POST("/v1/accounts", CreateAccount)
	ts := httptest.NewServer(server)
	ts.Close()

	t.Run("User Created", func(t *testing.T) {
		params := CreateAccountRequest{
			Email:          "test@test.com",
			Password:       "12345",
			RepeatPassword: "12345",
		}

		body, _ := json.Marshal(params)

		apitest.New().
			Handler(server).
			Post("/v1/accounts").
			JSON(string(body)).
			Expect(t).
			Assert(jsonpath.Equal(`$.data.email`, params.Email)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Params not valid", func(t *testing.T) {
		apitest.New().
			Handler(server).
			Post("/v1/accounts").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
}
