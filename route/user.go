package route

import (
	"core.yuyuid.id/internal/app/controller"
	"core.yuyuid.id/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func APIUserV1(app fiber.Router) {
	private := app.Group("user", middleware.EnsureAuthToken(true))

	private.Get("", controller.APIUserList)
}
