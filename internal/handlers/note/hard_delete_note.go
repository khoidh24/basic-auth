package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HardDeleteNote(c *fiber.Ctx) error {
	// Get the note ID from the request parameters
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
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

	// Create the filter for the hard delete
	filter := bson.M{
		"_id":    objID,
		"userId": user.ID,
	}

	// Perform the hard delete
	res, err := mgm.Coll(&features.Note{}).DeleteOne(c.Context(), filter)
	if err != nil || res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Note not found or unauthorized",
		})
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Note permanently deleted",
	})
}
