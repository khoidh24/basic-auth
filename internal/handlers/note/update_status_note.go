package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToggleActiveNote(c *fiber.Ctx) error {
	// Get note ID from params
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	// Get note from database
	note := &features.Note{}
	findErr := mgm.Coll(note).FindByID(objID, note)
	if findErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Note not found",
		})
	}

	// Get user from context
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil || note.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Update note status
	type payload struct {
		IsActive bool `json:"isActive"`
	}
	var body payload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request body",
		})
	}

	// Update note status
	note.IsActive = body.IsActive
	if err := mgm.Coll(note).Update(note); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Note updated successfully",
	})
}
