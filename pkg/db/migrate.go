package db

import (
	"log"

	"gorm.io/gorm"
)

func RunMigrate(g *gorm.DB) {
	err := g.AutoMigrate(Models...)

	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err.Error())
		return
	}

	log.Println("Migrate successfully")
}
