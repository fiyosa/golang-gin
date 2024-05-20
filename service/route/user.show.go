package route

import (
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/hash"
	"go-gin/pkg/helper"
	"go-gin/service/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 	Get user
// @Description Get user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "id"
// @Success 	200 {object} dto.UserShowResponse "ok"
// @Security	BearerAuth
// @Router 		/user/{id} [get]
func UserShow(c *gin.Context) {
	userAuth := db.User{}
	if !userAuth.GetUser(c) {
		return
	}

	user_id := c.Param("id")
	getId, _ := hash.Decode(user_id)

	user := db.User{}
	user.Show(getId)

	if user.Id == 0 {
		helper.SendError(
			c,
			lang.L(lang.SetL().NOT_FOUND, gin.H{"operator": lang.SetL().USER}),
			http.StatusNotFound,
		)
		return
	}

	id, _ := hash.Encode(user.Id)
	helper.SendData(
		c,
		lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, gin.H{"operator": lang.SetL().USER}),
		dto.UserShowDataResponse{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
	)
}
