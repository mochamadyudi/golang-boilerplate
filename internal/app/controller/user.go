package controller

import (
	"core.yuyuid.id/internal/app/model"
	"core.yuyuid.id/internal/app/repository"
	"core.yuyuid.id/internal/app/service"
	"core.yuyuid.id/internal/database"
	"core.yuyuid.id/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func APIUserList(c *fiber.Ctx) error {
	db := database.DB
	pagination := utils.Paginate(c)
	var total int64

	srv := service.NewUserService(repository.UserRepositoryImpl{DB: database.DB})

	if err := db.Model(&model.User{}).Count(&total).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("Failed to get data", &utils.Response[any]{}))
	}

	users, err := srv.Repo.FindAll(c.Context(), func(db *gorm.DB) *gorm.DB {
		db = db.Limit(pagination.Limit).Offset(pagination.Offset)
		return db
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("Failed to fetch users", &utils.Response[any]{}))
	}
	// if err := db.Model(&model.User{}).Limit(pagination.Limit).Offset(pagination.Offset).Find(&users).Error; err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError("Failed to fetch users", &utils.Response[any]{}))
	// }

	totalPages := int64(0)
	if total > 0 {
		totalPages = total / int64(pagination.Limit)
		if total%int64(pagination.Limit) > 0 {
			totalPages++
		}
	}
	// var data []model.UserResponse
	// for _, user := range users {
	// 	data = append(data, model.UserResponse{
	// 		ID:              user.ID,
	// 		FirstName:       user.FirstName,
	// 		LastName:        user.LastName,
	// 		Email:           user.Email,
	// 		EmailVerifiedAt: user.EmailVerifiedAt,
	// 	})
	// }

	return c.JSON(utils.ResponseSuccess("Successfully", &utils.Response[[]model.User]{
		Data: users,
		Pagination: &utils.Pagination{
			Page:    pagination.Page,
			Limit:   pagination.Limit,
			Total:   int(total),
			Maxpage: int(totalPages),
		},
	}))
}
