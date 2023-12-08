package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iconza98/trivia-api/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
