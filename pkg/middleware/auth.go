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

		_, err := jwt.Verify(c, token)

		if err != nil {
			helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
			return
		}

		user := db.User{}

		db.G.Where(&db.User{Token: token}).First(&user)

		if user.Id == 0 {
			helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
