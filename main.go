package main

import (
	"alirah/app"
	"alirah/database"
)

func main() {
	// database
	database.Connect()
	database.Migrate()

	app.StartApp()
}
