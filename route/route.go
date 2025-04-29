package route

import "github.com/gofiber/fiber/v2"

func Routes(app fiber.Router) {
	// registere route in here

	api := app.Group("/api/v1")
	APIAuthV1(api)
}
