package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Url string `required:"true"`
}

func Load() Config {
	var cfg Config

	envconfig.MustProcess("", &cfg)

	return cfg
}
