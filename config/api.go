package config

import "github.com/caarlos0/env/v6"

type API struct {
	HTTPPort int `env:"HTTP_PORT" envDefault:"8080"`
	DB       Database
}

func NewAPI() (API, error) {
	c := API{}
	err := env.Parse(&c)
	return c, err
}
