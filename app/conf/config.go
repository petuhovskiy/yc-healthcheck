package conf

import (
	"github.com/caarlos0/env/v6"
)

type App struct {
	Server      Server
	Postgres    Postgres
	Healthcheck Healthcheck
}

func ParseEnv() (*App, error) {
	cfg := App{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
