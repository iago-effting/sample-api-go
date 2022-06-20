package handlers

import (
	"encoding/json"
	"fmt"
	"iago-effting/api-example/pkg/accounts"
	"iago-effting/api-example/pkg/authentication"
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

func TestMe(t *testing.T) {
	ClearTable("users")

	server := gin.Default()
	server.POST("/v1/me", ViewAccount)

	ts := httptest.NewServer(server)
	ts.Close()

	t.Run("Get profile", func(t *testing.T) {
		var jwtService = authentication.JWTAuthService()

		params := accounts.User{
			Email:    "test@test.com",
			Password: "12345",
		}

		FactoryCreateAccount(params)
		token := jwtService.GenerateToken(params.Email)

		apitest.New().
			Handler(server).
			Post("/v1/me").
			Header("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect(t).
			Assert(jsonpath.Equal(`$.data.email`, params.Email)).
			//Assert(jsonpath.JWTHeaderEqual(fromAuthHeader, `$.alg`, "HS256")).
			Status(http.StatusOK).
			End()
	})
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
