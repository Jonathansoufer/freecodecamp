package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/Jonathansoufer/go-crm-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/Jonathansoufer/go-crm-fiber/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.CreateLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}

func initDatabase(){
	database.ConnectDB()
	db = database.GetDB()
	db.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	database.CloseDB()
}