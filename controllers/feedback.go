package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
	"strconv"
)

func AddFeedback(c *fiber.Ctx) error {
	db := database.DB
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["description"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "description is required!",
		})
	}
	feedback := models.Feedbacks{
		Description: data["description"],
	}
	if err := db.Create(&feedback).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error",
		})
	}
	return c.JSON(feedback)
}
func GetFeedbacks(c *fiber.Ctx) error {
	db := database.DB
	var feedbacks []models.Feedbacks
	db.Find(&feedbacks)
	return c.JSON(feedbacks)
}
