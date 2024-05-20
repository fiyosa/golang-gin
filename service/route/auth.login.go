package route

import (
	"errors"
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/hash"
	"go-gin/pkg/helper"
	"go-gin/pkg/jwt"
	"go-gin/service/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary 	Login
// @Description Login
// @Tags 		Auth
// @Accept 		json
// @Produce 	json
// @Param		payload body dto.AuthLoginPayload true "payload"
// @Success 	200 {object} dto.AuthLoginResponse "ok"
// @Router 		/auth/login [post]
func AuthLogin(c *gin.Context) {
	validated := dto.AuthLoginPayload{}

	if helper.Validate(c, &validated) {
		return
	}

	user := db.User{}

	if err := db.G.Where(&db.User{Username: validated.Username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendError(c, lang.L(lang.SetL().AUTH_NOT_FOUND, nil), http.StatusNotFound)
			return
		}
		helper.SendError(c, lang.L(lang.SetL().AUTH_FAILED, nil))
		return
	}

	if !hash.Verify(validated.Password, user.Password) {
		helper.SendError(c, lang.L(lang.SetL().AUTH_FAILED, nil))
		return
	}

	token, err := jwt.Create(helper.Int2Str(user.Id))

	if err != nil {
		helper.SendError(c, lang.L(err.Error(), nil))
		return
	}

	user.Token = token

	user.Update()

	c.JSON(200, gin.H{
		"token": user.Token,
	})
}
