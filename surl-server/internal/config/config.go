package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Get() *Config {
	godotenv.Load()
	return &Config{
		Port: ":" + os.Getenv("SERVER_PORT"),
	}
}
