package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	var envFile string

	switch env {
	case "dev":
		envFile = ".env.dev"
	case "prod":
		envFile = ".env.prod"
	default:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("⚠️  Не удалось загрузить %s файл: %v", envFile, err)
	}
}
