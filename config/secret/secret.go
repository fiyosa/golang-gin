package secret

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var APP_ENV string
var PORT string
var APP_DEBUG string

func Setup() bool {
	status := true
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err.Error())
		status = false
	}

	APP_ENV = getEnv("APP_ENV", "development")
	PORT = getEnv("PORT", "4000")
	APP_DEBUG = getEnv("APP_DEBUG", "true")

	return status
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(1, key, value, ok)
		return value
	}
	return fallback
}
