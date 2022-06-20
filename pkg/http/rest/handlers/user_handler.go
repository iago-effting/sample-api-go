package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Createad!",
	})
}
