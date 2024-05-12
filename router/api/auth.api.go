package api

import (
	"gin/config/secret"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthSecret(c *gin.Context) {
	c.JSON(200, gin.H{
		"data-manual-test": gin.H{
			"PORT":      secret.GetEnv1("PORT", "hehe 1"),
			"APP_ENV":   secret.GetEnv1("APP_ENV", "hehe 2"),
			"APP_DEBUG": secret.GetEnv1("APP_DEBUG", "hehe3"),
		},
		"data-manual": gin.H{
			"APP_PORT":  os.Getenv("PORT"),
			"APP_ENV":   os.Getenv("APP_ENV"),
			"APP_DEBUG": os.Getenv("APP_DEBUG"),
		},
		"data": gin.H{
			"APP_PORT":  secret.APP_PORT,
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
