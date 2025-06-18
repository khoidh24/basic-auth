package note

import (
	features "leanGo/internal/models/features"

	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var body struct {
	IDs []string `json:"ids"`
}

func SoftDeleteManyNotes(c *fiber.Ctx) error {
	// Check if the request body is valid
	if err := c.BodyParser(&body); err != nil || len(body.IDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request or empty ID list",
		})
	}

	// Check if the user is authenticated
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Convert string IDs to ObjectIDs
	objectIDs := []primitive.ObjectID{}
	for _, id := range body.IDs {
		objID, convertObjectIdErr := primitive.ObjectIDFromHex(id)
		if convertObjectIdErr == nil {
			objectIDs = append(objectIDs, objID)
		}
	}

	// Check if the notes exist and belong to the user
	filter := bson.M{
		"_id":      bson.M{"$in": objectIDs},
		"userId":   user.ID,
		"isActive": true,
	}

	// Soft delete the notes
	update := bson.M{
		"$set": bson.M{"isActive": false},
	}

	// Perform the update
	res, err := mgm.Coll(&features.Note{}).UpdateMany(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":       200,
		"message":      "Notes soft deleted",
		"deletedCount": res.ModifiedCount,
	})
}
