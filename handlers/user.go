package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/aselasperera/go-fiber-crud-api/dbConfig"
	"github.com/aselasperera/go-fiber-crud-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var timeout = 10 * time.Second

func GetUsers(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cursor, err := dbConfig.DATABASE.Collection("Users").Find(ctx, bson.M{})
	fmt.Print(cursor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to fetch users"})
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to fetch users"})
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var user models.User
	err = dbConfig.DATABASE.Collection("Users").FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to fetch user"})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := dbConfig.DATABASE.Collection("Users").InsertOne(ctx, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to create user"})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid user ID"})
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	update := bson.M{"$set": user}
	_, err = dbConfig.DATABASE.Collection("Users").UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to update user"})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err = dbConfig.DATABASE.Collection("Users").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete user"})
	}

	return c.SendStatus(200)
}
