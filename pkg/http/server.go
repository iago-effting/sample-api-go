package http

import (
	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http/rest"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

type service struct {
	Logger log.Logger
	Port   string
}

type Service interface {
	Run() error
}

func NewServerService(port string, logger log.Logger) Service {
	return &service{
		Logger: logger,
		Port:   port,
	}
}

func (s service) Run() error {
	if configs.Env.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(Logger(s.Logger))
	router = rest.Router(router)

	return router.Run(s.Port)
}
