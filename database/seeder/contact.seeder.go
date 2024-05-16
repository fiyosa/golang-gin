package seeder

import (
	"go-gin/database/model"

	"gorm.io/gorm"
)

func ContactSeeder(g *gorm.DB) {
	var users []*model.User

	contacts := []*model.Contact{
		{FirstName: "admin", LastName: "admin", Email: "admin@gmail.com", Phone: "081234567890"},
		{FirstName: "user", LastName: "user", Email: "user@gmail.com", Phone: "081234567890"},
	}

	g.Find(&users)

	for i, user := range users {
		contacts[i].UserId = user.Id
	}

	g.Create(&contacts)
}
