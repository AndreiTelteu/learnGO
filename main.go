package main

import (
	"learngo/app/http"
	routes "learngo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
)

func main() {
	engine := jet.New("./resources/views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	http.Kernel(app)
	routes.Web(app)
	app.Listen(":3000")
}
