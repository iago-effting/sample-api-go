package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	DebugLevel(c, "CreateUser", "Started")
	DebugLevel(c, "CreateUser", "Done")

	c.JSON(http.StatusOK, gin.H{
		"data": "Createad!",
	})
}
