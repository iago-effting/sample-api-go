package rest

import (
	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/http/rest/handlers"
)

func Router(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/v1")
	{
		v1.POST("/users", handlers.CreateUser)
	}

	return app
}
