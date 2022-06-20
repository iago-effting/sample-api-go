package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

func Logger(logger log.Logger, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Set("db", db)

		c.Next()
	}
}
