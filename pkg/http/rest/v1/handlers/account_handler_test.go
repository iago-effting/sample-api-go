package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func init() {
	Setup()
}

func TestCreateUser(t *testing.T) {
	ClearTable("users")

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
