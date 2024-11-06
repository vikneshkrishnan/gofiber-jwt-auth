package main

import (
	"gofiber-jwt-auth/database"
	"gofiber-jwt-auth/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Connect to the database
	database.Connect()

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
