package controller

import (
	"core.yuyuid.id/internal/database"
	"core.yuyuid.id/internal/model"
	"core.yuyuid.id/internal/request"
	"core.yuyuid.id/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(c *fiber.Ctx) error {
	db := database.DB
	body := c.Locals("body")
	req, ok := body.(request.AuthLoginRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("failed to cast request body", &utils.Response[any]{}))
	}

	var user model.User
	if err := db.Where("email = ? AND deleted_at IS NULL", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("Invalid Credentials", &utils.Response[any]{}))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("Invalid Credentials", &utils.Response[any]{}))
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("Could not login", &utils.Response[any]{}))
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("OK", &utils.Response[any]{
		Data: fiber.Map{"token": token},
	}))
}
