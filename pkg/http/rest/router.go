package rest

import (
	"github.com/gin-gonic/gin"
	"iago-effting/api-example/pkg/http/rest/v1/midlewares"

	"iago-effting/api-example/pkg/http/rest/v1/handlers"
)

func Router(app *gin.Engine) *gin.Engine {

	v1 := app.Group("/v1")
	{
		v1.GET("/me", midlewares.AuthorizeJWT(), handlers.ViewAccount)
		v1.GET("/accounts", handlers.IndexAccount)
		v1.DELETE("/accounts/:id", handlers.DeleteAccount)

		v1.POST("/accounts", handlers.CreateAccount)
		v1.POST("/auth", handlers.AuthHandler)
	}

	return app
}

//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
