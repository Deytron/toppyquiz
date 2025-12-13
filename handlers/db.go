package handlers

import (
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	// Check if Db exists, if not panic, as it should have been created before the container is run
	var err error
	Db, err = gorm.Open(gormlite.Open("tq.db"), &gorm.Config{})

	if err != nil {
		panic("GoRM : failed to connect database" + err.Error())
	}

	// Migrate the schema, as in create all the tables from structs
	models := []any{
		&Question{},
		&Theme{},
		&Team{},
		&ItemType{},
		&Item{},
		&Quiz{},
	}
	derr := Db.AutoMigrate(models...)
	if derr != nil {
		panic("GoRM : failed to migrate database")
	} else {
		println("GoRM : database migrated successfully")
	}
}
