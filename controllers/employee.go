package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"fmt"
	"main.go/database"
	"main.go/models"

)


func Getemployee(c *fiber.Ctx) error {
	db := database.DB
	var employee []models.Employee
	db.Find(&employee)
	
	return c.JSON(employee)
}

func Postemployee(c *fiber.Ctx) error {
		db := database.DB
		var employee models.Employee

		err := c.BodyParser(&employee); 
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
	
		
		
		db.Create(&employee)
		return c.JSON(employee)
	}

	func DeleteEmp(c *fiber.Ctx) error {
		db := database.DB
		id := c.Params("id")
		var employee models.Employee
		err := db.Find(&employee, "id = ?", id).Error
			if err != nil {
				return c.JSON(fiber.Map{
					"status":  "error",
					"message": "error in delete employee",
				})
			}
		db.Unscoped().Delete(&employee)
			return c.JSON(fiber.Map{
				"status":  "success",
				"message": "Employee data deleted",
			})	
	}

	func UpdateEmp(c *fiber.Ctx) error {
		db := database.DB
		var employee models.Employee
		id := c.Params("id")
			
			err := c.BodyParser(&employee)
			if err != nil {
				return c.JSON(fiber.Map{
					"status":  "error",
					"message": "Review your inputs",
					"data":    err,
				})
			}
			employee.Id,err = strconv.Atoi(id)
			if err != nil {
				return c.JSON(fiber.Map{
					"status":  "error",
					"message": "Review your inputs",
					"data":    err,
				})
			}
			fmt.Println(err)
	
			db.Save(&employee)
				return c.JSON(fiber.Map{
				"status":  "success",
				"message": "employee found",
				"error": err,
				"data":  employee,
			})
	}
	