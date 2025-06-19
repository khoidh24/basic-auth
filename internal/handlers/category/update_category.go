package category

import (
	features "leanGo/internal/models/features"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCategoryInfo(c *fiber.Ctx) error {
	// Get ID
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	// Get category
	category := &features.Category{}
	if err := mgm.Coll(category).FindByID(objID, category); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Category not found",
		})
	}

	// Get request body
	type Body struct {
		Name        string `json:"categoryName"`
		Description string `json:"description"`
	}
	var body Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	// Update category
	if body.Name != "" {
		category.Name = body.Name
	}
	category.Description = body.Description
	category.UpdatedAt = time.Now().UTC()
	if err := mgm.Coll(category).Update(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Category updated successfully",
	})
}
