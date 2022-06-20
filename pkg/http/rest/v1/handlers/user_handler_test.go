package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/storage/database"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/stretchr/testify/assert"
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

type ResponseCreatedUser struct {
	Data struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}
}

func TestCreateUser(t *testing.T) {
	clearTables()

	var target *ResponseCreatedUser
	var expectedResponse string

	server := gin.Default()
	server.POST("/v1/users", CreateUser)

	ts := httptest.NewServer(server)
	defer ts.Close()

	// TODO: Maybe a factory?
	params := CreateUserRequest{
		Email:          "test@test.com",
		Password:       "12345",
		RepeatPassword: "12345",
	}

	url := fmt.Sprintf("%s/v1/users", ts.URL)
	body, _ := json.Marshal(params)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	defer response.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	json.NewDecoder(response.Body).Decode(&target)

	assert.Equal(t, params.Email, target.Data.Email)
	assert.IsType(t, expectedResponse, target.Data.Id)
}

func TestCreateUserWithNoParams(t *testing.T) {
	clearTables()

	server := gin.Default()
	server.POST("/v1/users", CreateUser)

	ts := httptest.NewServer(server)
	defer ts.Close()

	missing_params := CreateUserRequest{
		Email:    "test@test.com",
		Password: "12345",
	}

	body, _ := json.Marshal(missing_params)
	url := fmt.Sprintf("%s/v1/users", ts.URL)

	// TODO: we need a better response
	expected := `{"error":"Key: 'CreateUserRequest.RepeatPassword' Error:Field validation for 'RepeatPassword' failed on the 'required' tag"}`

	response, _ := http.Post(url, "application/json", bytes.NewBuffer(body))
	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	actual, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, expected, string(actual))
}
