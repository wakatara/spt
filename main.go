package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/pug"
	"github.com/wakatara/spt/database"
	"github.com/wakatara/spt/species"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	app.Settings.Templates = pug.New("./views", ".pug")
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "species.db")
	if err != nil {
		panic("Failed to connect to DB. 😢")
	}
	fmt.Println("DB connect successful.")

	database.DBConn.AutoMigrate(&species.Species{})
	fmt.Println("Database Species table migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) { c.Render("index", fiber.Map{"Title": "Hello World"}) })
	app.Get("/api/v1/species", species.GetAllSpecies)
	app.Get("/api/v1/species/:id", species.GetSpecies)
	app.Post("/api/v1/species", species.NewSpecies)
	app.Delete("/api/v1/species/:id", species.DeleteSpecies)
}
