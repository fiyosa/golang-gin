package route

import (
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/hash"
	"go-gin/pkg/helper"
	"go-gin/service/dto"

	"github.com/gin-gonic/gin"
)

// @Summary 	Register
// @Description Register
// @Tags 		Auth
// @Accept 		json
// @Produce 	json
// @Param		Payload  body dto.AuthRegisterPayload true "Payload"
// @Success 	200 {object} dto.AuthRegisterResponse "ok"
// @Router 		/auth/register [post]
func AuthRegister(c *gin.Context) {
	validated := dto.AuthRegisterPayload{}

	if helper.Validate(c, &validated) {
		return
	}

	user := db.User{}

	db.G.Where(&db.User{Username: validated.Username}).First(&user)

	if user.Id != 0 {
		helper.SendError(
			c,
			lang.L(lang.SetL().ALREADY_EXIST, gin.H{"operator": "Username"}),
		)
		return
	}

	newPassword, err := hash.Create(validated.Password)

	if err != nil {
		helper.SendError(c, err.Error())
		return
	}

	user.Username = validated.Username
	user.Name = validated.Name
	user.Password = newPassword

	user.Create()

	c.JSON(200, gin.H{
		"data": gin.H{
			"id":         user.Id,
			"username":   user.Username,
			"name":       user.Name,
			"pass":       user.Password,
			"created_at": helper.Time2Str(user.CreatedAt),
			"updated_at": helper.Time2Str(user.UpdatedAt),
		},
		"message": lang.L(lang.SetL().SAVED_SUCCESSFULLY, gin.H{"operator": lang.SetL().USER}),
	})
}
