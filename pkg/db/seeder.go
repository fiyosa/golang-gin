package db

import (
	"go-gin/database/seeder"

	"gorm.io/gorm"
)

func RunSeeder(g *gorm.DB) {
	seeder.UserSeeder(g)
	seeder.ContactSeeder(g)
	seeder.AddressSeeder(g)
}
