package route

import (
	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello User",
	})
}
