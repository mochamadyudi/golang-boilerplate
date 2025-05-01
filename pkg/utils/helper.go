package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetLocalUserID(c *fiber.Ctx) (uint, error) {
	raw := c.Locals("userID")
	if raw == nil {
		return 0, errors.New("User ID Not found")
	}
	userID, ok := raw.(uint)
	if !ok {
		return 0, errors.New("User ID not number")
	}

	return userID, nil
}
