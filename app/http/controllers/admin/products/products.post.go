package products

import (
	"learngo/app/helpers"
	"learngo/app/helpers/logger"
	"learngo/app/models"
	"learngo/database"

	"github.com/gofiber/fiber/v2"
)

type PostProductRequest struct {
	Name         string                     `validate:"required,min=3,max=32"`
	Price        float32                    `validate:"omitempty,number"`
	RegularPrice float32                    `validate:"omitempty,number"`
	SalePrice    float32                    `validate:"omitempty,number"`
	Description  string                     `validate:"required,min=3,max=6000"`
	Categories   []ProductCategoriesRequest `validate:"omitempty,dive"`
}
type ProductCategoriesRequest struct {
	Id   int    `validate:"omitempty,number"`
	Name string `validate:"omitempty"`
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

	var categories []models.Category
	if product.Categories != nil && len(product.Categories) > 0 {
		for _, category := range product.Categories {
			if category.Id == 0 && len(category.Name) > 0 {
				// Create new category
				newCategory := models.Category{
					Name: category.Name,
				}
				database.DB().Create(&newCategory)
				categories = append(categories, newCategory)
			}
			if category.Id > 0 {
				// Find existing category
				var existingCategory models.Category
				result := database.DB().First(&existingCategory, category.Id)
				if result.RowsAffected > 0 {
					categories = append(categories, existingCategory)
				}
			}
		}
	}

	logger.Debug("Final cat list", categories)

	//Return user
	return c.JSON(categories)
}
