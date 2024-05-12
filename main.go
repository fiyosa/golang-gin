package main

import (
	"gin/config/secret"
	"gin/router"
)

func main() {
	if !secret.Setup() {
		// panic("Error loading .env file")
	}

	r := router.Setup()

	r.Run(":" + secret.PORT)
}
