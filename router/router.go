package router

import (
	docs "go-gin/docs"
	"go-gin/pkg/middleware"
	"go-gin/pkg/secret"
	"go-gin/service/route"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := configRouter()
	p := r.Group("/api") // prefix /api

	p.POST("/auth/register", route.AuthRegister)
	p.POST("/auth/login", route.AuthLogin)
	p.GET("/auth/secret", route.AuthSecret)

	p.Use(middleware.Auth())
	{
		p.GET("/auth/user", route.AuthUser)
		p.GET("/user", route.UserIndex)
		p.GET("/user/:id", route.UserShow)
	}

	r.NoRoute(route.AuthNotFound)

	return r
}

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

	r.LoadHTMLGlob("service/view/*.html")
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "welcome.html", gin.H{}) })

	docs.SwaggerInfo.Host = secret.APP_URL
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DocExpansion("none"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))
	r.GET("/swagger", func(c *gin.Context) { c.Redirect(http.StatusFound, "/swagger/index.html") })

	return r
}
