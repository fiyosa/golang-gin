package api

import "github.com/gin-gonic/gin"

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
