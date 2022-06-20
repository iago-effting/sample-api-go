package configs

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
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
		DSN string `env:"DATABASE_DRIVER"`
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

	_, filename, _, _ := runtime.Caller(0)
	baseFolder := filepath.Dir(filename)

	commonFileName := fmt.Sprintf("%s/%s.toml", baseFolder, environment)
	envFileName := fmt.Sprintf("%s/common.toml", baseFolder)

	commomFeeder := feeder.Toml{Path: commonFileName}
	tomlFeeder := feeder.Toml{Path: envFileName}
	envFeeder := feeder.Env{}

	c := config.New()

	c.AddFeeder(commomFeeder)
	c.AddFeeder(tomlFeeder)
	c.AddFeeder(envFeeder)

	c.AddStruct(&appconfiguration)

	err := c.Feed()
	if err != nil {
		fmt.Println("Error", err.Error())
		level.Error(s.Logger).Log(err)
	}

	Env = appconfiguration

	fmt.Println("appconfiguration", appconfiguration.Database.DSN)

	level.Debug(s.Logger).Log("LoadEnvVar", true)
}
