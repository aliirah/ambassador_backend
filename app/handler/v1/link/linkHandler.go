package link

import (
	"alirah/app/domain"
	"alirah/app/resource/link"
	"alirah/database"
	"alirah/util/rest"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Index(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))

	var links []domain.Link
	database.DB.Where("user_id = ?", userId).Find(&links)

	return rest.Ok(c, fiber.Map{
		"links": link.Collection(&links),
	})
}
