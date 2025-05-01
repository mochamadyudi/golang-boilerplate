package middleware

import (
	"log"
	"strings"

	"core.yuyuid.id/internal/app/repository"
	"core.yuyuid.id/internal/app/service"
	"core.yuyuid.id/internal/database"
	"core.yuyuid.id/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func EnsureAuthToken(enabled bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		var userID uint = 0

		if header != "" {
			parts := strings.Split(header, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				valToken := parts[1]

				claims, err := utils.ParseToken(valToken)

				if err == nil {
					userIDFloat, ok := claims["user_id"].(float64)
					if !ok {
						log.Fatalln("user_id is not valid number")
					}
					userID = uint(userIDFloat)

					repo := repository.UserRepositoryImpl{DB: database.DB}
					userService := service.NewUserService(repo)

					err := userService.IsExist(c.Context(), userID)
					if err != nil {
						userID = 0
					}
				}
			}
		}
		if userID == 0 {
			if enabled {
				return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError("Unauthorized: valid token required", &utils.Response[any]{
					Meta: &utils.Meta{
						Code: "401",
					},
				}))
			}
		}

		if enabled {
			c.Locals("userID", userID)
		}
		return c.Next()
	}
}
