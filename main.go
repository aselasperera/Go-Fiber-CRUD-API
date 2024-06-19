package main

import (
	"github.com/aselasperera/go-fiber-crud-api/dbConfig"
	"github.com/aselasperera/go-fiber-crud-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	dbConfig.ConnectToMongoDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Go Fiber CRUD API!")
	})

	routes.SetupUserRoutes(app)

	app.Listen(":3000")
}
