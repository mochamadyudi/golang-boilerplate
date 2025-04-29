package controller

import (
	"core.yuyuid.id/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess("OK", &utils.Response[any]{}))
}
