package route

import (
	"go-gin/lang"

	"github.com/gin-gonic/gin"
)

func AuthNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": lang.L(lang.SetL().API_NOT_FOUND, nil),
	})
}
