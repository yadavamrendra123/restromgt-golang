package database

import (
	"gorm.io/gorm"
	"log"
	"restro-mgt/models"
)

func RunMigrations(db *gorm.DB) {

	if db == nil {
		log.Fatalf("database connection is nil")
		return
	}

	// Run migrations
	err := db.AutoMigrate(
		&models.Restaurant{},
		&models.Event{},
		&models.TimeEntry{},
		&models.Secret{},
		&models.CustomFormatModel{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
