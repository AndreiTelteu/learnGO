package main

import (
	"learngo/app/http"
	"learngo/database"
	routes "learngo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
	"github.com/joho/godotenv"
)

func main() {

	// Load env vars from file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Init jet templating engine
	engine := jet.New("./resources/views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Connect database and do migrations
	database.Connect()
	database.Migrate()

	// Init routes and middleware
	http.Kernel(app)
	routes.Web(app)

	// Start / Listen
	app.Listen(":3000")
}
