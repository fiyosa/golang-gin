package middleware

import (
	"go-gin/lang"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil)},
		)
		c.Next()
	}
}
