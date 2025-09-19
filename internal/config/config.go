package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Url                  string `required:"true"`
	GoogleServiceAccount string `split_words:"true" required:"true"`
	SpreadsheetId        string `split_words:"true" required:"true"`
	SheetName            string `split_words:"true" default:"Sheet1"`
}

func Load() Config {
	var cfg Config

	envconfig.MustProcess("", &cfg)

	return cfg
}
