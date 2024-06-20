package main

import (
	"log"

	"github.com/aselasperera/go-fiber-crud-api/dbConfig"
	"github.com/aselasperera/go-fiber-crud-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Enable CORS for all origins
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	dbConfig.ConnectToMongoDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Go Fiber CRUD API!")
	})

	routes.SetupUserRoutes(app)

	log.Fatal(app.Listen(":3001")) // Changed port to 3001
}
