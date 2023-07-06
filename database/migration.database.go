package database

import (
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
	// &models.Role{},
	)
}
