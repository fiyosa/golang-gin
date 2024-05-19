package middleware

import (
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/helper"
	"go-gin/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken := c.GetHeader("Authorization")
		token := strings.Split(getToken, " ")[1]

		if getToken == "" && token == "" {
			helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
			return
		}

		result, err := jwt.Verify(c, token)

		if err != nil {
			helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
			return
		}

		user := db.User{}
		user.Show(helper.Str2Int(result))
		c.Set("user", user)
		c.Next()
	}
}
