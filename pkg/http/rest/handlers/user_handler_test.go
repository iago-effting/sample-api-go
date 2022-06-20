package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
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
	expected := `{"data":"Createad!"}`

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	defer response.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	actual, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, expected, string(actual))
}

func TestCreateUserWithNoParams(t *testing.T) {
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
