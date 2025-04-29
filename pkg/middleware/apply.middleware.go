package middleware

import (
	"fmt"
	"time"

	"core.yuyuid.id/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func ApplyMiddleware(app *fiber.App) {
	cnf := config.Get()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("https://yuyuid.id, http://localhost, http://localhost:%s", cnf.Server.Port),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowCredentials: true,
		MaxAge:           36800,
	}))
	app.Use(compress.New())

	// implement limiter
	app.Use(limiter.New(limiter.Config{
		Max:               5,
		Expiration:        5 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}
