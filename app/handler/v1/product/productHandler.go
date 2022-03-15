package product

import (
	domain "alirah/app/domain/product"
	productRequest "alirah/app/request/v1/product"
	productResource "alirah/app/resource/product"
	"alirah/database"
	"alirah/util/rest"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Index(c *fiber.Ctx) error {
	var products []domain.Product
	database.DB.Find(&products)

	return rest.Ok(c, fiber.Map{
		"products": productResource.Collection(&products),
	})
}

func Store(c *fiber.Ctx) error {
	var body productRequest.StoreData

	if err := c.BodyParser(&body); err != nil {
		return rest.BadRequest(c, err)
	}

	if err := productRequest.StoreValidate(&body); err != nil {
		return rest.ValidationError(c, err)
	}

	// TODO handle store image
	product := domain.Product{
		Title:       body.Title,
		Description: body.Description,
		Image:       body.Image,
		Price:       body.Price,
	}
	database.DB.Create(&product)

	return rest.Ok(c, fiber.Map{
		"product": productResource.SingleResource(&product),
	})
}

func Show(c *fiber.Ctx) error {
	var product domain.Product
	id, _ := strconv.Atoi(c.Params("id"))

	res := database.DB.
		Where("id = ?", id).
		Find(&product)

	if res.RowsAffected == 0 {
		return rest.NotFound(c)
	}

	return rest.Ok(c, fiber.Map{
		"product": productResource.SingleResource(&product),
	})
}
