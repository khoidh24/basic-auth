package category

import (
	features "leanGo/internal/models/features"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCategory(c *fiber.Ctx) error {
	// Parse ID
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	// Find category by ID
	category := &features.Category{}
	if err := mgm.Coll(category).FindByID(objID, category); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Category not found",
		})
	}

	type Body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsActive    bool   `json:"isActive"`
	}

	var body Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	// Update category
	category.Name = body.Name
	category.Description = body.Description
	category.IsActive = body.IsActive

	if err := mgm.Coll(category).Update(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to update category",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Update category successfully",
	})
}
