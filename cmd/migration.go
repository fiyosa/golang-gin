package cmd

import (
	"go-gin/pkg/db"
	"go-gin/pkg/seeder"
	"log"

	"gorm.io/gorm"
)

func RunSeeder(g *gorm.DB) {
	seeder.UserSeeder(g)
	seeder.ContactSeeder(g)
	seeder.AddressSeeder(g)
}

func RunMigrate(g *gorm.DB) {
	err := g.AutoMigrate(db.Models...)

	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err.Error())
		return
	}

	log.Println("Migrate successfully")
}

func RunDropAllTable(g *gorm.DB) {
	err := g.Migrator().DropTable(db.Models...)

	if err != nil {
		log.Fatalf("Failed drop all table: %v", err.Error())
		return
	}

	log.Println("Deleted all table successfully")
}
