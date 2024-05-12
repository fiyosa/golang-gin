package secret

import (
	"os"

	"github.com/joho/godotenv"
)

var APP_ENV string
var PORT string
var APP_DEBUG string

func Setup() bool {
	if err := godotenv.Load(".env"); err != nil {
		return false
	}

	APP_ENV = getEnv("APP_ENV", "development")
	PORT = getEnv("PORT", "4000")
	APP_DEBUG = getEnv("APP_DEBUG", "true")

	return true
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnv1(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
