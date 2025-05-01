package handler

import (
	"time"

	"core.yuyuid.id/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ResponseMapUserLoader struct {
	ID              uint               `json:"id"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Email           string             `json:"email"`
	EmailVerifiedAt *time.Time         `json:"email_verified_at"`
	Status          string             `json:"status"`
	CreatedAt       *time.Time         `json:"created_at"`
	UpdatedAt       *time.Time         `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt    `json:"deleted_at"`
	Role            *fiber.Map         `json:"role"`
	Permissions     []model.Permission `json:"permissions"`
}

func MapUserLoader(user model.User) ResponseMapUserLoader {
	var role *fiber.Map
	if user.RoleUser.Role.ID != 0 {
		tmp := fiber.Map{
			"id":    user.RoleUser.Role.ID,
			"name":  user.RoleUser.Role.Name,
			"slug":  user.RoleUser.Role.Slug,
			"level": user.RoleUser.Role.Level,
		}
		role = &tmp
	}

	return ResponseMapUserLoader{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		Status:          user.Status,
		CreatedAt:       &user.CreatedAt,
		UpdatedAt:       &user.UpdatedAt,
		DeletedAt:       &user.DeletedAt,
		Role:            role,
		Permissions:     user.RoleUser.Role.Permissions,
	}
}
