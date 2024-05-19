package main

import (
	"fmt"
	"go-gin/cmd"
	_ "go-gin/docs"
	"go-gin/pkg/db"
	"go-gin/pkg/secret"
	"go-gin/pkg/validation"
	"go-gin/router"
)

// @title Tag Service API
// @version 1.0
// @description A Tag Service API in Go useing Gin Framework
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	secret.Setup()

	db.Setup()

	if cmd.Setup() {
		return // using the option
	}

	if err := validation.Setup(); err != nil {
		fmt.Println(err.Error())
		return
	}

	r := router.Setup()

	r.Run(":" + secret.PORT)
}
