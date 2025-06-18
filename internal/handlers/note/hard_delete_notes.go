package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HardDeleteManyNotes(c *fiber.Ctx) error {
	// Define the request body structure
	var body struct {
		IDs []string `json:"ids"`
	}

	// Parse the request body
	if err := c.BodyParser(&body); err != nil || len(body.IDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request or empty ID list",
		})
	}

	// Get the user from the context

	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Convert the string IDs to ObjectIDs
	objectIDs := []primitive.ObjectID{}
	for _, id := range body.IDs {
		objID, convertObjectIdErr := primitive.ObjectIDFromHex(id)
		if convertObjectIdErr == nil {
			objectIDs = append(objectIDs, objID)
		}
	}

	// Create the filter for the hard delete
	filter := bson.M{
		"_id":    bson.M{"$in": objectIDs},
		"userId": user.ID,
	}

	// Perform the hard delete
	res, err := mgm.Coll(&features.Note{}).DeleteMany(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":       200,
		"message":      "Notes permanently deleted",
		"deletedCount": res.DeletedCount,
	})
}
