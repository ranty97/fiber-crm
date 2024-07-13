package main

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", GetLeads)
	app.Post("api/v1/lead", CreateLead)
	app.Delete("api/v1/lead/:id", DeleteLead)
	app.Get("api/v1/lead/:id", GetLead)
}

func initDB() {

}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen("localhost:9091")
}
