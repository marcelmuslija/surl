package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPass string
}

func Get() *Config {
	godotenv.Load()
	return &Config{
		Port:   ":" + os.Getenv("SERVER_PORT"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
	}
}
