package configs

import (
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type ConfigEnv struct {
	Server struct {
		Port int `env:"SERVER_PORT"`
	}
	Debug bool   `env:"DEBUG"`
	Name  string `env:"ENV"`
}

type service struct {
	Env    string
	Logger log.Logger
}

type Service interface {
	LoadEnvVars()
}

var Env ConfigEnv

func NewConfigService(env string, logger log.Logger) Service {
	return &service{
		Env:    env,
		Logger: logger,
	}
}

func (s service) LoadEnvVars() {
	environment := "dev"
	if env := s.Env; env != "" {
		environment = env
	}

	appconfiguration := ConfigEnv{}

	fileName := fmt.Sprintf("configs/%s.toml", environment)

	commomFeeder := feeder.Toml{Path: "configs/common.toml"}
	tomlFeeder := feeder.Toml{Path: fileName}
	envFeeder := feeder.Env{}

	c := config.New()

	c.AddFeeder(commomFeeder)
	c.AddFeeder(tomlFeeder)
	c.AddFeeder(envFeeder)

	c.AddStruct(&appconfiguration)

	err := c.Feed()
	if err != nil {
		level.Error(s.Logger).Log(err)
	}

	Env = appconfiguration

	level.Debug(s.Logger).Log("LoadEnvVar", true)
}
