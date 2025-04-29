package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UUID            string         `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Email           string         `json:"email,omitempty" gorm:"unique"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	Password        string         `json:"-"`
	RememberToken   string         `json:"remember_token,omitempty"`
	Status          string         `json:"status,omitempty"`
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
