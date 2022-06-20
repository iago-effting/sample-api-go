package midlewares

import (
	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/logs"
)

func Logger(logger logs.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)

		c.Next()
	}
}
