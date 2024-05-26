package handlers

import (
	"github.com/aselasperera/go-fiber-crud-api/database"
	"github.com/aselasperera/go-fiber-crud-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID"})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID"})
	}
	database.DB.Delete(&user)
	return c.SendStatus(200)
}
