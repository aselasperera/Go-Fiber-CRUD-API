package main

import (
	"github.com/aselasperera/go-fiber-crud-api/database"
	"github.com/aselasperera/go-fiber-crud-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDatabase()

	// Define a root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Go Fiber CRUD API!")
	})

	routes.SetupUserRoutes(app)

	app.Listen(":3000")
}
