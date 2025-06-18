package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToggleActiveNotes(c *fiber.Ctx) error {
	type RequestBody struct {
		IDs      []string `json:"ids"`
		IsActive bool     `json:"isActive"` // Used for both soft delete and restore
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil || len(body.IDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request or empty ID list",
		})
	}

	// Get authenticated user
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Convert string IDs to ObjectIDs
	objectIDs := make([]primitive.ObjectID, 0, len(body.IDs))
	for _, id := range body.IDs {
		objID, convertErr := primitive.ObjectIDFromHex(id)
		if convertErr == nil {
			objectIDs = append(objectIDs, objID)
		}
	}

	if len(objectIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "No valid note IDs provided",
		})
	}

	// Filter notes by ownership and target IDs
	filter := bson.M{
		"_id":    bson.M{"$in": objectIDs},
		"userId": user.ID,
	}

	// Update isActive according to request
	update := bson.M{
		"$set": bson.M{
			"isActive": body.IsActive,
		},
	}

	// Perform update
	res, err := mgm.Coll(&features.Note{}).UpdateMany(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to update notes",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       200,
		"message":      "Notes updated successfully",
		"updatedCount": res.ModifiedCount,
	})
}
