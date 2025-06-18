package category

import (
	features "leanGo/internal/models/features"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeactiveManyCategory(c *fiber.Ctx) error {
	// Parse request body
	var payload struct {
		Ids []string `json:"ids"`
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
	for _, id := range payload.Ids {
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

	// Update categories
	_, err := mgm.Coll(&features.Category{}).UpdateMany(c.Context(),
		bson.M{"_id": bson.M{"$in": objIDs}},
		bson.M{"$set": bson.M{"isActive": false}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Update categories successfully",
	})
}
