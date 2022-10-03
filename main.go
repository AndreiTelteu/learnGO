package main

import (
	"learngo/app/http"
	routes "learngo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	http.Kernel(app)
	routes.Web(app)
	app.Listen(":3000")
}
