package rest

import (
	"iago-effting/api-example/pkg/http/rest/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/v1")
	{
		v1.GET("/account", handlers.IndexAccount)
		v1.POST("/account", handlers.CreateAccount)
		v1.GET("/account/:id", handlers.ViewAccount)
		v1.DELETE("/account/:id", handlers.DeleteAccount)
	}

	return app
}
