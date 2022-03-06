package admin

import (
	v1 "alirah/app/handler/v1"
	"github.com/gofiber/fiber/v2"
)

func MapUrl(r fiber.Router) {
	r.Post("/register", v1.Register)
}
