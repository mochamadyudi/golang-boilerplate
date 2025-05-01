package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"not null;nique"`
	Level       int64          `json:"level" gorm:"not null"`
	CreatedAt   *time.Time     `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Users       []User         `json:"users,omitempty" gorm:"many2many:user_roles"`
	Permissions []Permission   `json:"permissions,omitempty" gorm:"many2many:role_permissions"`
}

func (Role) TableName() string {
	return "roles"
}
