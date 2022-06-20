package http

import (
	"iago-effting/go-template/pkg/http/rest"

	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()
	router = rest.Router(router)

	return router.Run(":8080")
}
