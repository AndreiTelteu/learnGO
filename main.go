package main

import (
	routes "learngo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Web(app)
	app.Listen(":3000")
}
