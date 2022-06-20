package http

import (
	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http/rest"

	"github.com/gin-gonic/gin"
)

func Run(port string) error {
	if configs.Env.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router = rest.Router(router)

	return router.Run(port)
}
