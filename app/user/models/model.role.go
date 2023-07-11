package user

import (
	role "go-echo-modular/app/role/models"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"index:idx_name,unique"`
	Password  string         `json:"-"`
	Role      role.Role      `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	RoleID    int            `json:"role_id"`
	IsActive  bool           `json:"is_active,omitempty" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
