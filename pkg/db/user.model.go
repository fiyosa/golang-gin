package db

import (
	"go-gin/lang"
	"go-gin/pkg/helper"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Token     string    `json:"token" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	Contacts []Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *User) GetUser(c *gin.Context) bool {
	user, exists := c.Get("user")
	if !exists {
		helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
		return false
	}
	userObj, ok := user.(User)
	if !ok {
		helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
		return false
	}
	*u = userObj
	return true
}

func (u *User) Create() *gorm.DB {
	return G.Create(&u)
}

func (u *User) Show(id int) *gorm.DB {
	return G.First(&u, id)
}

func (u *User) Update() *gorm.DB {
	return G.Save(&u)
}

func (u *User) Delete(id int) *gorm.DB {
	return G.Delete(&u, id)
}
