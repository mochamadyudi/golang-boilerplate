package model

import (
	"context"
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
	RoleUser        *RoleUser      `gorm:"foreignKey:UserId" json:"role_user,omitempty"`
}

func (User) TableName() string {
	return "users"
}

type UserService interface {
	List(ctx context.Context) ([]User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	SaveUser(ctx context.Context, user *User) error
	IsUserExist(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, user *User) error
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByUuid(ctx context.Context, uuid string) (User, error)
	FindById(ctx context.Context, id uint) (User, error)
	IsExist(ctx context.Context, id uint) error
	Save(ctx context.Context, c *User) error
	Update(ctx context.Context, c *User) error
	Delete(ctx context.Context, id uint) error
}
