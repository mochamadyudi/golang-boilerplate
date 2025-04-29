package main

import (
	"fmt"

	"core.yuyuid.id/internal/database"
	ws "core.yuyuid.id/internal/websocket"
	"core.yuyuid.id/pkg/config"
	"core.yuyuid.id/pkg/middleware"
	"core.yuyuid.id/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	fibCnf := config.FiberConf()
	app := fiber.New(fibCnf)

	middleware.ApplyMiddleware(app)

	database.Connect(cnf.Database)
	ws.Register(app)

	route.Routes(app)

	app.Listen(fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port))

}
