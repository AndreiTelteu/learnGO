package home

import (
	"fmt"
	"learngo/app/http"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	sess := http.GetSession(c)
	defer sess.Save()

	sess.Set("User", fiber.Map{"Name": "guy"})

	seuser := sess.Get("User").(fiber.Map)

	return c.SendString(fmt.Sprintf("Hello, World %v 👋!", seuser["Name"]))
}
