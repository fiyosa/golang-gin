package cmd

import (
	"flag"
	"go-gin/pkg/db"
)

func Setup() bool {
	dropFlag := flag.Bool("drop", false, "Drop the database tables")
	seedFlag := flag.Bool("seed", false, "Seed the database with initial data")
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")

	flag.Parse()
	status := false

	if *dropFlag {
		RunDropAllTable(db.G)
		status = true
	}

	if *migrateFlag {
		RunMigrate(db.G)
		status = true
	}

	if *seedFlag {
		RunSeeder(db.G)
		status = true
	}

	return status
}
