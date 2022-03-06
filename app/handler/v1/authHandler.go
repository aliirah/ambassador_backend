package v1

import (
	"alirah/app/request/v1/auth"
	"alirah/util/rest"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var body auth.RegisterData
	if err := c.BodyParser(&body); err != nil {
		return rest.BadRequest(c, err)
	}

	err := auth.Validate(&body)
	if err != nil {
		return rest.ValidationError(c, err)
	}

	return rest.Ok(c, map[string]string{
		"message": "Hello World!",
	})
}
