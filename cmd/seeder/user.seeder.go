package seeder

import (
	"go-gin/pkg/db"
	"go-gin/pkg/hash"
	"go-gin/pkg/helper"
	"go-gin/pkg/jwt"

	"gorm.io/gorm"
)

func UserSeeder(g *gorm.DB) {
	password, _ := hash.Create("Password")

	users := []*db.User{
		{Username: "admin", Name: "Admin", Password: password},
		{Username: "user", Name: "User", Password: password},
	}

	g.Create(&users)

	for _, user := range users {
		if result, err := jwt.Create(helper.Int2Str(int(user.Id))); err != nil {
			user.Token = result
		}
	}

	g.Updates(&users)
}
