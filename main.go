package main

import (
	"alirah/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// database
	database.Connect()
	database.Migrate()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	err := app.Listen(":8000")
	if err != nil {
		log.Fatalln("error serving")
	}
}
