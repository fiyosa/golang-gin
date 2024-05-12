package router

import (
	"gin/config/secret"
	"gin/pkg/helper"
	"gin/router/api"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	if value, _ := helper.Str2Bool(secret.APP_DEBUG); value {
		r.Use(gin.Logger())
	}

	p := r.Group("/api") // prefix

	auth := p.Group("/auth")
	{
		auth.GET("/secret", api.AuthSecret)
		auth.GET("/login", api.AuthLogin)
		auth.GET("/register", api.AuthRegister)
	}

	user := p.Group("/user")
	{
		user.GET("/login", api.AuthLogin)
		user.GET("/register", api.AuthRegister)
	}

	return r
}
