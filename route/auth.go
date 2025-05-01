package route

import (
	"core.yuyuid.id/internal/app/controller"
	"core.yuyuid.id/internal/app/request"
	"core.yuyuid.id/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func APIAuthV1(route fiber.Router) {
	api := route.Group("auth")
	api.Post("/login", middleware.ValidatorMiddleware[request.AuthLoginRequest](), controller.AuthLogin)
	api.Post("/register", middleware.ValidatorMiddleware[request.AuthRegisterRequest](), controller.AuthRegister)

	private := api.Group("", middleware.EnsureAuthToken(true))
	private.Get("/user", controller.AuthLoadUser)
}
