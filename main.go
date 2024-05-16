package main

import (
	"go-gin/cmd"
	"go-gin/pkg/db"
	"go-gin/pkg/secret"
	"go-gin/router"
)

func main() {
	secret.Setup()

	db.Setup()

	if cmd.Setup() {
		return // using the option
	}

	r := router.Setup()

	r.Run(":" + secret.PORT)
}
