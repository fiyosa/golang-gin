package api

import (
	"gin/config/secret"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthSecret(c *gin.Context) {
	c.JSON(200, gin.H{
		"data-manual": gin.H{
			"PORT":      os.Getenv("PORT"),
			"APP_ENV":   os.Getenv("APP_ENV"),
			"APP_DEBUG": os.Getenv("APP_DEBUG"),
		},
		"data": gin.H{
			"APP_PORT":  secret.PORT,
			"APP_ENV":   secret.APP_ENV,
			"APP_DEBUG": secret.APP_DEBUG,
		},
		"message": "Secret retrieved sucessfully",
	})
}

func AuthLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func AuthRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
