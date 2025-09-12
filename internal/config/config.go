package config

import (
	"os"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
)

func NewConfig() *helper.Config {
	cfg := &helper.Config{
		Prefix:         os.Getenv("PREFIX"),
		BaseUrl:        os.Getenv("BASE_URL"),
		UserServiceUrl: os.Getenv("USER_SERVICE_URL"),
	}

	if cfg.BaseUrl == "" {
		panic("BASE_URL not found")
	}

	if cfg.UserServiceUrl == "" {
		panic("USER_SERVICE_URL not found")
	}

	return cfg
}
