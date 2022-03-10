package admin

import (
	v1 "alirah/app/handler/v1"
	"github.com/gofiber/fiber/v2"
)

func MapUrl(r fiber.Router) {
	r.Post("/register", v1.Register)
	r.Post("/login", v1.Login)
	r.Get("/user", v1.User)
	r.Post("/logout", v1.Logout)
}
