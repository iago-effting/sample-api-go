package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type ConfigEnv struct {
	Migrations struct {
		Dir string
	}
	Server struct {
		Port int `env:"SERVER_PORT"`
	}
	Database struct {
		DSN      string `env:"DATABASE_DSN"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Host     string `env:"DATABASE_HOST"`
		Port     int    `env:"DATABASE_PORT"`
		Name     string `env:"DATABASE_NAME"`
	}
	Debug struct {
		Database    bool `env:"DEBUG_DATABASE"`
		Application bool `env:"DEBUG_APPLICATION"`
	}
	Name string `env:"ENV"`
}

type service struct {
	Env    string
	Logger *logrus.Logger
}

type Service interface {
	LoadEnvVars()
}

var Env ConfigEnv

func NewConfigService(env string, logger *logrus.Logger) Service {
	return &service{
		Env:    env,
		Logger: logger,
	}
}

func (s service) LoadEnvVars() {
	appConfiguration := ConfigEnv{}

	_, filename, _, _ := runtime.Caller(0)
	baseFolder := filepath.Dir(filename)

	commonFileName := fmt.Sprintf("%s/%s.toml", baseFolder, s.Env)
	envFileName := fmt.Sprintf("%s/common.toml", baseFolder)

	commonFeeder := feeder.Toml{Path: commonFileName}
	tomlFeeder := feeder.Toml{Path: envFileName}
	envFeeder := feeder.Env{}

	c := config.New()

	c.AddFeeder(commonFeeder)
	c.AddFeeder(tomlFeeder)
	c.AddFeeder(envFeeder)

	c.AddStruct(&appConfiguration)

	err := c.Feed()
	if err != nil {
		s.Logger.Error(err)
	}

	Env = appConfiguration

	s.Logger.Debug("LoadEnvVar", true)
}
