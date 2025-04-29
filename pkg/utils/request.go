package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetBody[T any](c *fiber.Ctx, rules T) (T, error) {

	body := c.Body()
	// Unmarshal body ke struct req
	if err := json.Unmarshal([]byte(body), &rules); err != nil {
		return rules, err
	}
	return rules, nil
}
