package main

import (
	"gin/config/secret"
	"gin/router"
	"os"
)

func main() {
	secret.Setup()
	// if !secret.Setup() {
	// panic("Error loading .env file")
	// }

	r := router.Setup()

	r.Run(":" + os.Getenv("APP_PORT"))
}
