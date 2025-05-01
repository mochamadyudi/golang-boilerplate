package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug" gorm:"uniqueIndex:permissions_slug_unique;not null"`
	Description string         `json:"description,omitempty"`
	Model       string         `json:"model"`
	CreatedAt   *time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Role        []Role         `gorm:"many2many:role_permissions" json:"role,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
