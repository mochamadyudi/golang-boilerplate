package config

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConf() fiber.Config {

	return fiber.Config{
		BodyLimit:     40 * 1024 * 1024,
		Network:       "tcp",
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   time.Second * time.Duration(10),
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	}
}
