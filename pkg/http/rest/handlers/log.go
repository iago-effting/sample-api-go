package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func DebugLevel(c *gin.Context, key string, value interface{}) {
	logger, hasLogger := c.Get("logger")

	if !hasLogger {
		return
	}

	level.Debug((logger).(log.Logger)).Log(key, value)
}
