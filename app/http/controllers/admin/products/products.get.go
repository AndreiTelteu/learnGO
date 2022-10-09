package products

import (
	"learngo/app/helpers"

	"github.com/gofiber/fiber/v2"
)

type GetProductRequest struct {
	Name  string `query:"name" validate:"omitempty,min=3,max=32"`
	Email string `validate:"omitempty,email,min=6,max=32"`
}

func Get(c *fiber.Ctx) error {
	filters := new(GetProductRequest)
	if err := c.QueryParser(filters); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := helpers.Validate(*filters)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"errors":  errors,
		})
	}

	return c.JSON(filters)
}
