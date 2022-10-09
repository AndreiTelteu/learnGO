package products

import (
	"learngo/app/helpers"

	"github.com/gofiber/fiber/v2"
)

type PostProductRequest struct {
	Name  string `validate:"required,min=3,max=32"`
	Email string `validate:"required,email,min=6,max=32"`
}

func Post(c *fiber.Ctx) error {
	product := new(PostProductRequest)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := helpers.Validate(*product)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"errors":  errors,
		})
	}
	//Return user
	return c.JSON(product)
}
