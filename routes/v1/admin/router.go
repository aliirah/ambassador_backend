package admin

import (
	ambassadorHandler "alirah/app/handler/v1/ambassador"
	authHandler "alirah/app/handler/v1/auth"
	linkHandler "alirah/app/handler/v1/link"
	orderHandler "alirah/app/handler/v1/order"
	productHandler "alirah/app/handler/v1/product"
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

	auth.Get("/ambassadors", ambassadorHandler.Index)

	auth.Get("/products", productHandler.Index)
	auth.Post("/products", productHandler.Store)
	auth.Get("/products/:id", productHandler.Show)
	auth.Put("/products/:id", productHandler.Update)
	auth.Delete("/products/:id", productHandler.Delete)

	auth.Get("/users/:id/links", linkHandler.Index)

	auth.Get("/orders", orderHandler.Index)
}
