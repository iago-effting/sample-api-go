package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"iago-effting/api-example/pkg/accounts"
	"iago-effting/api-example/pkg/authentication"
	accountDb "iago-effting/api-example/pkg/storage/database/account"
)

func init() {
	Setup()
}

func TestAuthUser(t *testing.T) {
	server := gin.Default()
	server.POST("/v1/authentication", AuthHandler)
	ts := httptest.NewServer(server)

	defer ts.Close()
	defer ClearTable("users")

	ClearTable("users")

	t.Run("User authenticate", func(t *testing.T) {
		account := accounts.User{
			Email:    "example@test.com",
			Password: "123456789",
		}

		FactoryCreateAccount(account)

		params := authentication.Credentials{
			Email:    account.Email,
			Password: account.Password,
		}

		body, _ := json.Marshal(params)

		apitest.New().
			Handler(server).
			Post("/v1/authentication").
			JSON(body).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Invalid credentials", func(t *testing.T) {
		params := authentication.Credentials{
			Email:    "invalid",
			Password: "invalid",
		}

		body, _ := json.Marshal(params)

		apitest.New().
			Handler(server).
			Post("/v1/authentication").
			JSON(body).
			Expect(t).
			Assert(jsonpath.Equal(`$.message`, "Credentials are not valid")).
			Status(http.StatusForbidden).
			End()
	})
}

func FactoryCreateAccount(userParams accounts.User) *accounts.User {
	var ctx = context.Background()
	var usersRepository accounts.Repository = accountDb.Repo()

	user, _ := usersRepository.Save(ctx, userParams)

	return user
}
