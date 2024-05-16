package route

import (
	"go-gin/lang"

	"github.com/gin-gonic/gin"
)

func AuthLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": lang.L(
			lang.SetL().RETRIEVED_SUCCESSFULLY,
			gin.H{"operator": lang.SetL().USER}),
	})
}
