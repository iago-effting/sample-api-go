package rest

import (
	"iago-effting/api-example/pkg/http/rest/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/v1")
	{
		v1.GET("/users", handlers.IndexUser)
		v1.POST("/users", handlers.CreateUser)
		v1.GET("/users/:id", handlers.ViewUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)
	}

	return app
}
