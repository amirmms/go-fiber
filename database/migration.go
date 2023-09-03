package database

import (
	"fiber_test/entities/models"
	"gorm.io/gorm"
	"log"
)

func Migration(DB *gorm.DB) {
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
