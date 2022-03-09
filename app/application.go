package app

import (
	"alirah/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func StartApp() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	
	// routes
	routes.Setup(app)

	// listen to port
	appPort := os.Getenv("APP_PORT")
	err = app.Listen(":" + appPort)
	if err != nil {
		log.Fatalln("error serving")
	}
}
