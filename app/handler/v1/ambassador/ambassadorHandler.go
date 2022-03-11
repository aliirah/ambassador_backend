package ambassador

import (
	"alirah/app/domain"
	"alirah/app/resource/user"
	"alirah/database"
	"alirah/util/rest"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var ambassadors []domain.User
	database.DB.
		Where("is_ambassador = true").
		Find(&ambassadors)

	return rest.Ok(c, fiber.Map{
		"ambassadors": user.Collection(&ambassadors),
	})
}
