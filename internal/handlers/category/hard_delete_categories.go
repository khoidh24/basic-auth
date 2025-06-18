package category

import (
	features "leanGo/internal/models/features"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HardDeleteManyCategory(c *fiber.Ctx) error {
	// Parse request body
	var payload struct {
		IDs []string `json:"ids"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	// Convert IDs to ObjectID
	objIDs := make([]primitive.ObjectID, 0)
	invalidIDs := 0
	for _, id := range payload.IDs {
		if objID, err := primitive.ObjectIDFromHex(id); err == nil {
			objIDs = append(objIDs, objID)
		} else {
			invalidIDs++
		}
	}

	// Check if there are any valid IDs
	if len(objIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "No valid IDs provided",
		})
	}

	// Delete categories
	_, err := mgm.Coll(&features.Category{}).DeleteMany(c.Context(), bson.M{"_id": bson.M{"$in": objIDs}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to permanently delete categories",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Permanently delete categories successfully",
	})
}
