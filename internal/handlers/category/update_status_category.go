package category

import (
	features "leanGo/internal/models/features"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToggleActiveCategory(c *fiber.Ctx) error {
	// Get ID
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	// Get category by ID
	category := &features.Category{}
	if err := mgm.Coll(category).FindByID(objID, category); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Category not found",
		})
	}

	// Get payload isActive
	type Body struct {
		IsActive bool `json:"isActive"`
	}

	var body Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	// Set isActive to category
	category.IsActive = body.IsActive
	if err := mgm.Coll(category).Update(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Category active state updated",
	})
}
