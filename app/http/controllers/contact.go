package contact

import "github.com/gofiber/fiber/v2"

func Contact(c *fiber.Ctx) error {

	return c.SendString("Hello, World contact 👋!")
}
