package rest

import (
	"iago-effting/api-example/pkg/http/rest/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/v1")
	{
		v1.GET("/accounts", handlers.IndexAccount)
		v1.POST("/accounts", handlers.CreateAccount)
		v1.GET("/accounts/:id", handlers.ViewAccount)
		v1.DELETE("/accounts/:id", handlers.DeleteAccount)
	}

	return app
}
