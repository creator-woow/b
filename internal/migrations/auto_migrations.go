package migrations

import (
	"gorm.io/gorm"
	"happy-server/internal/models"
	"log"
)

func AutoMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Panicln("Some of auto migrations failed: " + err.Error())
	}
}
