package admin

import (
	authHandler "alirah/app/handler/v1/auth"
	"alirah/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapUrl(r fiber.Router) {
	// guest
	r.Post("/register", middleware.IsGuest, authHandler.Register)
	r.Post("/login", middleware.IsGuest, authHandler.Login)

	// authenticated user
	auth := r.Use(middleware.IsAuth)
	auth.Get("/user", authHandler.User)
	auth.Post("/logout", authHandler.Logout)
	auth.Put("/update", authHandler.UpdateInfo)
	auth.Put("/update/password", authHandler.UpdatePassword)
}
