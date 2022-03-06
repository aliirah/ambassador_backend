package rest

import (
	"github.com/gofiber/fiber/v2"
)

func Ok(c *fiber.Ctx, data interface{}) error {
	err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ok",
		"data":    data,
	})
	if err != nil {
		return err
	}
	return nil
}

func BadRequest(c *fiber.Ctx, message interface{}) error {
	if message == "" {
		message = "Bad Request"
	}
	err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
	if err != nil {
		return err
	}
	return nil
}

func ValidationError(c *fiber.Ctx, data map[string]string) error {
	err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Validation Error",
		"errors":  data,
	})
	if err != nil {
		return err
	}
	return nil
}

func NotFound(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not found",
	})
	if err != nil {
		return err
	}
	return nil
}
