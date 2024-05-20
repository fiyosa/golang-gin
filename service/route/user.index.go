package route

import (
	"fmt"
	"go-gin/lang"
	"go-gin/pkg/db"
	"go-gin/pkg/hash"
	"go-gin/pkg/helper"
	"go-gin/service/dto"

	"github.com/gin-gonic/gin"
)

// @Summary 	Get users
// @Description Get users
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		page query int false "page"
// @Param 		limit query int false "limit"
// @Param 		keyword query string false "keyword"
// @Param 		orderBy query string false "orderBy"
// @Param 		sortedBy query string false "sortedBy"
// @Success 	200 {object} dto.UserIndexResponse "ok"
// @Security	BearerAuth
// @Router 		/user [get]
func UserIndex(c *gin.Context) {
	user := db.User{}
	if !user.GetUser(c) {
		return
	}

	queryStr := helper.QueryStr(c)

	fmt.Println(queryStr)

	users := []db.User{}
	db.G.
		Offset(helper.Offset(queryStr.Page, queryStr.Limit)).
		Limit(queryStr.Limit).
		Order(queryStr.OrderBy+" "+queryStr.SortedBy).
		Where("username like ?", "%"+queryStr.Keyword+"%").
		Or("name like ?", "%"+queryStr.Keyword+"%").
		Find(&users)

	var countUsers int64
	db.G.
		Model(&db.User{}).
		Where("username like ?", "%"+queryStr.Keyword+"%").
		Or("name like ?", "%"+queryStr.Keyword+"%").
		Count(&countUsers)

	newUsers := []dto.UserIndexDataResponse{}

	for _, v := range users {
		id, _ := hash.Encode(v.Id)
		newUsers = append(newUsers, dto.UserIndexDataResponse{
			Id:        id,
			Username:  v.Username,
			Name:      v.Name,
			CreatedAt: helper.Time2Str(v.CreatedAt),
			UpdatedAt: helper.Time2Str(v.UpdatedAt),
		})
	}

	helper.SendDatas(
		c,
		lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, gin.H{"operator": lang.SetL().USER}),
		newUsers,
		helper.Paginate{
			Page:  queryStr.Page,
			Limit: queryStr.Limit,
			Total: int(countUsers),
		},
	)
}
