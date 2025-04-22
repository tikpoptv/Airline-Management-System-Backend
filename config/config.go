package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var requiredKeys = []string{
	"DB_DSN",
	"PORT",
	"JWT_SECRET",
	"LOG_LEVEL",
}

func LoadEnv() {
	needDotenv := false

	for _, key := range requiredKeys {
		if os.Getenv(key) == "" {
			needDotenv = true
			break
		}
	}

	if needDotenv {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️ Warning: .env file not found or could not be loaded")
		} else {
			log.Println("✅ Loaded .env file (fallback for missing keys)")
		}
	}

	for _, key := range requiredKeys {
		if os.Getenv(key) == "" {
			log.Printf("⚠️ Warning: Required environment variable %s is still missing\n", key)
		}
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
