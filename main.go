package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"main.go/database"
	"main.go/routes"
)

func main(){
	app := fiber.New()
	database.Connect()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*",
	}))
	routes.EmployeeRoute(app)
	log.Fatal(app.Listen(":8080"))
}