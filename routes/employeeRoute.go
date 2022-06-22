package routes

import(
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
)

func EmployeeRoute(app *fiber.App){
	app.Get("/getEmployee",controllers.Getemployee)
	app.Post("/createEmployee",controllers.Postemployee)
	app.Delete("/deleteEmployee/:id",controllers.DeleteEmp)
	app.Put("/updateEmployee/:id",controllers.UpdateEmp)

}