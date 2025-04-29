package route

import (
	"core.yuyuid.id/internal/controller"
	"core.yuyuid.id/internal/request"
	"core.yuyuid.id/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func APIAuthV1(route fiber.Router) {
	api := route.Group("auth")
	api.Post("/login", middleware.ValidatorMiddleware(&request.AuthLoginRequest{}), controller.AuthLogin)
}
