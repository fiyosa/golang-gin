package route

import (
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/hash"
	"go-gin/pkg/helper"
	"go-gin/service/dto"

	"github.com/gin-gonic/gin"
)

// @Summary 	Get user by auth
// @Description Get user by auth
// @Tags 		Auth
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.AuthUserResponse "ok"
// @Security	BearerAuth
// @Router 		/auth/user [get]
func AuthUser(c *gin.Context) {
	user := db.User{}
	if !user.GetUser(c) {
		return
	}

	id, _ := hash.Encode(user.Id)
	helper.SendData(
		c,
		lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, gin.H{"operator": lang.SetL().USER}),
		dto.AuthUserDataResponse{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
	)
}
