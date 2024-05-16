package db

import (
	"log"

	"gorm.io/gorm"
)

func RunDropAllTable(g *gorm.DB) {
	err := g.Migrator().DropTable(Models...)

	if err != nil {
		log.Fatalf("Failed drop all table: %v", err.Error())
		return
	}

	log.Println("Deleted all table successfully")
}
