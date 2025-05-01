package controller

import (
	"core.yuyuid.id/internal/app/handler"
	"core.yuyuid.id/internal/app/model"
	"core.yuyuid.id/internal/app/repository"
	"core.yuyuid.id/internal/app/request"
	"core.yuyuid.id/internal/app/service"
	"core.yuyuid.id/internal/database"
	"core.yuyuid.id/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func AuthRegister(c *fiber.Ctx) error {
	db := database.DB
	body := c.Locals("body")
	req, ok := body.(request.AuthRegisterRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("failed to cast request body", &utils.Response[any]{}))
	}
	repo := repository.UserRepositoryImpl{DB: db}
	service := service.NewUserService(repo)

	if err := service.Repo.IsExistEmail(c.Context(), string(req.Email)); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("Email has already registered", &utils.Response[any]{}))
	}

	// hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 4)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("Failed to has password", &utils.Response[any]{}))
	}

	user := model.User{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Status:    "INACTIVE",
	}

	user.Password = string(hashedPassword)

	if err := service.SaveUser(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("Error Saving User", &utils.Response[any]{}))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ResponseSuccess(
		"User Registered successfully",
		&utils.Response[model.User]{
			Data: user,
		}))
}

func AuthLoadUser(c *fiber.Ctx) error {
	userID, err := utils.GetLocalUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError(err.Error(), &utils.Response[any]{}))
	}

	srv := service.NewUserService(repository.UserRepositoryImpl{DB: database.DB})

	data, err := srv.FindByID(c.Context(), userID, func(db *gorm.DB) *gorm.DB {
		return db.Preload("RoleUser.Role.Permissions")
	})

	if err != nil && data.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("Unauthorized!", &utils.Response[any]{}))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.ResponseSuccess(
		"Successfully",
		&utils.Response[any]{
			Data: handler.MapUserLoader(data),
		}))
}
