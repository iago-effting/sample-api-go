package rest

import (
	"iago-effting/api-example/pkg/http/rest/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/v1")
	{
		v1.POST("/users", handlers.CreateUser)
	}

	return app
}
