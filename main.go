package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ranty97/fiber-crm-basic/database"
	"github.com/ranty97/fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.Delete)
	app.Get("/api/v1/lead/:id", lead.GetLead)
}

func initDB() {
	var err error
	database.DBconnection, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed connect to database")
	}
	fmt.Println("Opened connection to database")
	database.DBconnection.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDB()
	setupRoutes(app)
	app.Listen("localhost:9091")
	defer database.DBconnection.Close()
}
