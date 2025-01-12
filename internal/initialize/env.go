package initialize

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"happy-server/internal/config"
	"log"
)

func Env() *config.EnvConfig {
	var c config.EnvConfig
	if loadErr := godotenv.Load(); loadErr != nil {
		log.Panicln("Error loading .env file")
	}
	parseErr := env.Parse(&c)
	if parseErr != nil {
		log.Panicln(parseErr)
	}
	return &c
}
