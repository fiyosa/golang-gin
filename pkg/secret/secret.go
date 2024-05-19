package secret

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	APP_ENV    string
	APP_LOCALE string
	APP_SECRET string
	APP_URL    string

	DB_HOST string
	DB_PORT string
	DB_NAME string
	DB_USER string
	DB_PASS string
)

func env() {
	PORT = getEnv("PORT", "4000")
	APP_ENV = getEnv("APP_ENV", "development")
	APP_LOCALE = getEnv("APP_LOCALE", "en")
	APP_SECRET = getEnv("APP_SECRET", "secret")
	APP_URL = getEnv("APP_URL", "http://localhost:4000")

	DB_HOST = getEnv("DB_HOST", "localhost")
	DB_PORT = getEnv("DB_PORT", "5432")
	DB_NAME = getEnv("DB_NAME", "golang-gin")
	DB_USER = getEnv("DB_USER", "postgres")
	DB_PASS = getEnv("DB_PASS", "")
}

func Setup() bool {
	status := true
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
		status = false
	}
	env()
	return status
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
