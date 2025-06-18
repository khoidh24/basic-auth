package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateNote(c *fiber.Ctx) error {
	// Get note ID from params
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	// Get existing note
	note := &features.Note{}
	findErr := mgm.Coll(note).FindByID(objID, note)
	if findErr != nil || !note.IsActive {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Note not found or inactive",
		})
	}

	// Check if user owns the note
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil || note.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Parse request body
	type payload struct {
		Title      string `json:"noteTitle"`
		Content    string `json:"content"`
		CoverImage string `json:"coverImage"`
	}
	var body payload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request body",
		})
	}

	// Validate request body
	if body.Title != "" {
		note.Title = body.Title
	}
	note.Content = body.Content
	note.CoverImage = body.CoverImage
	note.UpdatedAt = time.Now().UTC()

	// Update note in database
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
