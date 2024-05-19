package route

import (
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/helper"

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

	if err := user.GetUser(c); err != nil {
		helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
		return
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"id":         user.Id,
			"username":   user.Username,
			"name":       user.Name,
			"created_at": helper.Time2Str(user.CreatedAt),
			"updated_at": helper.Time2Str(user.UpdatedAt),
		},
		"message": lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, gin.H{"operator": lang.SetL().USER}),
	})
}
