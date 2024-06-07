package database

import (
	"gorm.io/gorm"
	"log"
	"restro-mgt/models" // Make sure to import the correct models package where your models are defined
)

func RunMigrations(db *gorm.DB) {
	// Ensure the DB is not nil
	if db == nil {
		log.Fatalf("database connection is nil")
		return
	}

	// Run migrations
	err := db.AutoMigrate(
		&models.Restaurant{}, // Replace with other models as necessary
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
