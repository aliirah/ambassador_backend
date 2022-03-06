package routes

import (
	v1Admin "alirah/routes/v1/admin"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	routesV1 := api.Group("/v1")

	routesV1Admin := routesV1.Group("/admin")
	v1Admin.MapUrl(routesV1Admin)
}
