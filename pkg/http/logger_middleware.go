package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

func Logger(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)

		c.Next()
	}
}
