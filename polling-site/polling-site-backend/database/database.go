package database

import (
	"log"
	"polling-site/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database 
func InitDB(){
	var err error
	DB, err = gorm.Open(sqlite.Open("polls.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// auto migrate the schema
	DB.AutoMigrate(&models.Poll{}, &models.Option{})
}


