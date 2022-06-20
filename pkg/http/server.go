package http

import (
	"iago-effting/api-example/pkg/http/rest"

	"github.com/gin-gonic/gin"
)

func Run(port string) error {
	router := gin.Default()
	router = rest.Router(router)

	return router.Run(port)
}
