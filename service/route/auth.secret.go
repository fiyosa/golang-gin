package route

import (
	"fmt"
	"go-gin/database/model"
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/jwt"
	"go-gin/pkg/secret"

	"github.com/gin-gonic/gin"
)

func AuthSecret(c *gin.Context) {
	rawJwt := "testing"
	jwtCreate, errJC := jwt.Create(rawJwt)
	if errJC != nil {
		c.JSON(400, gin.H{"message": jwtCreate})
		return
	}
	jwtVerify, errJV := jwt.Verify(jwtCreate, c)
	if errJV != nil {
		c.JSON(400, gin.H{"message": jwtCreate})
		return
	}

	var users []model.User
	result := db.G.Preload("Contacts.Addresses").Find(&users)

	if result.Error != nil {
		fmt.Println("Error result:", result.Error.Error())
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"PORT":       secret.PORT,
			"APP_ENV":    secret.APP_ENV,
			"APP_LOCALE": secret.APP_LOCALE,

			"DB_HOST": secret.DB_HOST,
			"DB_PORT": secret.DB_PORT,
			"DB_NAME": secret.DB_NAME,
			"DB_USER": secret.DB_USER,
			"DB_PASS": secret.DB_PASS,

			"JWT Raw":    rawJwt,
			"JWT Create": jwtCreate,
			"JWT Verify": jwtVerify,

			"user_id": users,
		},
		"message": lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, gin.H{"operator": "Secret"}),
	})
}
