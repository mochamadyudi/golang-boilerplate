package middleware

import (
	"fmt"

	"core.yuyuid.id/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Validate[T any](c *fiber.Ctx) (T, error) {
	var req T

	if err := c.BodyParser(&req); err != nil {
		return req, nil
	}

	err := validate.Struct(req)
	if err != nil {
		return req, err
	}
	return req, nil
}

func ValidatorMiddleware[T any]() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req T
		// Parse body ke struct rules
		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseError("Invalid Request", &utils.Response[any]{}))
		}

		// Validasi
		err := validate.Struct(req)
		if err != nil {
			var validationErrors []string
			for _, e := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, DefaultValidatorMessage(e)) // ambil nama field saja
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ResponseSuccess("Error Validation", &utils.Response[any]{Error: validationErrors}))
		}
		ctx.Locals("body", req)
		return ctx.Next()
	}
}

func DefaultValidatorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address.", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters.", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters.", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s.", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s.", e.Field(), e.Param())
	case "numeric":
		return fmt.Sprintf("%s must be a numeric value.", e.Field())
	case "url":
		return fmt.Sprintf("%s must be a valid URL.", e.Field())
	default:
		return fmt.Sprintf("%s %s", e.Field(), e.Param())
	}
}
