package home

import (
	"learngo/app/http"
	"learngo/app/models"
	"learngo/database"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	sess := http.GetSession(c)
	defer sess.Save()
	db := database.DB()

	var firstProduct models.Product
	db.First(&firstProduct)

	sess.Set("User", fiber.Map{"Name": "guy"})

	seuser := sess.Get("User").(fiber.Map)

	// return c.SendString(fmt.Sprintf("Hello, World %v 👋!", seuser["Name"]))
	return c.Render("home", fiber.Map{
		"Name":    seuser["Name"],
		"Product": firstProduct,
	})
}
