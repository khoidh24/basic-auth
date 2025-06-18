package category

import (
	features "leanGo/internal/models/features"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToggleActiveCategories(c *fiber.Ctx) error {
	// Parse request body
	var payload struct {
		Ids      []string `json:"ids"`
		IsActive bool     `json:"isActive"` // Value to apply
	}
	if err := c.BodyParser(&payload); err != nil || len(payload.Ids) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request or missing IDs",
		})
	}

	// Convert IDs to ObjectID
	objIDs := make([]primitive.ObjectID, 0, len(payload.Ids))
	for _, id := range payload.Ids {
		if objID, err := primitive.ObjectIDFromHex(id); err == nil {
			objIDs = append(objIDs, objID)
		}
	}

	if len(objIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "No valid category IDs provided",
		})
	}

	// Perform update
	update := bson.M{
		"$set": bson.M{
			"isActive": payload.IsActive,
		},
	}

	res, err := mgm.Coll(&features.Category{}).UpdateMany(c.Context(),
		bson.M{"_id": bson.M{"$in": objIDs}}, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to update categories",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       200,
		"message":      "Categories updated successfully",
		"updatedCount": res.ModifiedCount,
	})
}
