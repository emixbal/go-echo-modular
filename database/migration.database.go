package database

import (
	role "go-echo-modular/app/role/models"
	user "go-echo-modular/app/user/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&role.Role{},
		&user.User{},
	)
}
