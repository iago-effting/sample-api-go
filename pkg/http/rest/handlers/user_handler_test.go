package handlers

import (
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

	url := fmt.Sprintf("%s/v1/users", ts.URL)
	expected := `{"data":"Createad!"}`

	response, err := http.Post(url, "application/json", nil)
	defer response.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	actual, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, expected, string(actual))
}
