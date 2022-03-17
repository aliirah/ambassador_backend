package link

import (
	"alirah/app/domain"
	orderResource "alirah/app/resource/order"
	"alirah/database"
	"alirah/util/rest"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var orders []domain.Order
	database.DB.Find(&orders)

	return rest.Ok(c, fiber.Map{
		"orders": orderResource.Collection(&orders),
	})
}
