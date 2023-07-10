package database

import (
	role "go-echo-modular/app/role/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&role.Role{},
	)
}
