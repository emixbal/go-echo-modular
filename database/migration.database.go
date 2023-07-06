package database

import (
	"go-echo-modular/app/role"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&role.Model{},
	)
}
