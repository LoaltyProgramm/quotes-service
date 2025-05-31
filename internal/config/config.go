package config

import (
	"os"

	"github.com/LoaltyProgramm/quotes-service/internal/models/config"
)

func NewConfig() *config.Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &config.Config{
		Port: port,
	}
}
