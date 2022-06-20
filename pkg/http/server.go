package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http/rest"
)

type service struct {
	Logger *logrus.Logger
	Port   string
}

type Service interface {
	Run() error
}

func NewServerService(port string, logger *logrus.Logger) Service {
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
	router.Use(ginlogrus.Logger(s.Logger))
	router = rest.Router(router)

	return router.Run(s.Port)
}
