package web

import (
	contact "learngo/app/http/controllers"
	home "learngo/app/http/controllers/home"

	"github.com/gofiber/fiber/v2"
)

func Web(app *fiber.App) {
	app.Get("/", home.Homepage)
	app.Get("/contact", contact.Contact)
}
