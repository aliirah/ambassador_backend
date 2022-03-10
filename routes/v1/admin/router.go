package admin

import (
	v1 "alirah/app/handler/v1"
	"alirah/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapUrl(r fiber.Router) {
	// guest
	r.Post("/register", middleware.IsGuest, v1.Register)
	r.Post("/login", middleware.IsGuest, v1.Login)

	// authenticated user
	auth := r.Use(middleware.IsAuth)
	auth.Get("/user", v1.User)
	auth.Post("/logout", v1.Logout)
}
