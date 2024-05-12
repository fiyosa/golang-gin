package main

import (
	"gin/config/secret"
	"gin/router"
	"os"
)

func main() {
	secret.Setup()

	r := router.Setup()

	r.Run(":" + os.Getenv("PORT"))
}
