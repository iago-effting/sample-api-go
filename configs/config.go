package configs

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"

	"iago-effting/api-example/pkg/logs"
)

type ConfigEnv struct {
	Authentication struct {
		Secret  string `env:"AUTHENTICATION_SECRET"`
		Issuer  string `env:"AUTHENTICATION_ISSUER"`
		Expires int    `env:"AUTHENTICATION_HOURS_TO_EXPIRES"`
	}
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
	Logger logs.Logger
}

type Service interface {
	LoadEnvVars()
}

var Env ConfigEnv

func NewConfigService(env string, logger logs.Logger) Service {
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
