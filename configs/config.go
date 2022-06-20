package configs

import (
	"fmt"

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

var Env ConfigEnv

func SetVarEnvs(envMode string) {
	environment := "dev"
	if env := envMode; env != "" {
		environment = env
	}

	appconfiguration := ConfigEnv{}

	fileName := fmt.Sprintf("configs/%s.toml", environment)

	commomFeeder := feeder.Toml{Path: "configs/config.toml"}
	tomlFeeder := feeder.Toml{Path: fileName}
	envFeeder := feeder.Env{}

	c := config.New()

	c.AddFeeder(commomFeeder)
	c.AddFeeder(tomlFeeder)
	c.AddFeeder(envFeeder)

	c.AddStruct(&appconfiguration)

	err := c.Feed()
	if err != nil {
		panic(err)
	}

	Env = appconfiguration
}
