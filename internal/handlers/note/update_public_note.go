package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TogglePublicNote handles updating the isPublic field of a note.
func TogglePublicNote(c *fiber.Ctx) error {
	// Parse note ID
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid note ID",
		})
	}

	// Find note by ID
	note := &features.Note{}
	if existingNoteErr := mgm.Coll(note).FindByID(objID, note); existingNoteErr != nil || !note.IsActive {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Note not found or inactive",
		})
	}

	// Get user from context
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Check ownership
	if note.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  403,
			"message": "Forbidden",
		})
	}

	// Parse body
	type Body struct {
		IsPublic bool `json:"isPublic"`
	}

	var body Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	// Update isPublic and updatedAt
	note.IsPublic = body.IsPublic
	note.UpdatedAt = time.Now().UTC()

	if err := mgm.Coll(note).Update(note); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to update note public state",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Note public state updated successfully",
	})
}
