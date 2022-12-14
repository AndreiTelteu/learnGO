package web

import (
	contact "learngo/app/http/controllers"
	products "learngo/app/http/controllers/admin/products"
	home "learngo/app/http/controllers/home"

	"github.com/gofiber/fiber/v2"
)

func Web(app *fiber.App) {
	app.Get("/", home.Homepage)
	app.Get("/contact", contact.Contact)

	app.Get("/api/admin/products", products.Get)
	app.Post("/api/admin/products", products.Post)
}
