package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"iago-effting/api-example/pkg/http/rest/v1/midlewares"
	"iago-effting/api-example/pkg/logs"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http/rest"
)

type service struct {
	Logger logs.Logger
	Port   string
}

type Service interface {
	Run() error
}

func NewServerService(port string, logger logs.Logger) Service {
	return &service{
		Logger: logger,
		Port:   port,
	}
}

func (s service) Run() error {
	if !configs.Env.Debug.Application {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(midlewares.Logger(s.Logger))
	router.Use(ginlogrus.Logger(logrus.New()))

	router = rest.Router(router)

	return router.Run(s.Port)
}
