package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() {
	// Check if Db exists, if not panic, as it should have been created before the container is run
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("GoRM : failed to connect database")
	}

	// Migrate the schema, as in create all the tables from structs
	models := []any{
		&Question{},
		&Theme{},
		&Team{},
		&ItemType{},
		&Item{},
	}
	db.AutoMigrate(models...)
}
