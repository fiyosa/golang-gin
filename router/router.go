package router

import (
	"go-gin/pkg/middleware"
	"go-gin/pkg/secret"
	"go-gin/service/route"

	"github.com/gin-gonic/gin"
)

func configRouter() *gin.Engine {
	if secret.APP_ENV == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	if secret.APP_ENV == "development" {
		r.Use(gin.Logger())
	}
	r.SetTrustedProxies(nil)
	r.Use(middleware.Cors())

	return r
}

func Setup() *gin.Engine {
	r := configRouter()
	p := r.Group("/api") // prefix /api

	p.GET("/auth/login", route.AuthLogin)
	p.GET("/auth/secret", route.AuthSecret)

	p.Use(middleware.Auth())
	{
		p.GET("/user", route.UserIndex)
	}

	r.NoRoute(route.AuthNotFound)

	return r
}
