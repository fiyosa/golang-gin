package api

import (
	"gin/config/secret"

	"github.com/gin-gonic/gin"
)

func AuthSecret(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": gin.H{
			"PORT":      secret.PORT,
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
