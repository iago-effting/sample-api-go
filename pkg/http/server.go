package http

import (
	"database/sql"
	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http/rest"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

type service struct {
	Logger log.Logger
	Port   string
	DB     *sql.DB
}

type Service interface {
	Run() error
}

func NewServerService(port string, logger log.Logger, db *sql.DB) Service {
	return &service{
		Logger: logger,
		Port:   port,
		DB:     db,
	}
}

func (s service) Run() error {
	if !configs.Env.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(Logger(s.Logger, s.DB))
	router = rest.Router(router)

	return router.Run(s.Port)
}
