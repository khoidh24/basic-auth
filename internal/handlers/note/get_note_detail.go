package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNoteDetail(c *fiber.Ctx) error {
	noteID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	note := &features.Note{}
	if err := mgm.Coll(note).FindByID(objID, note); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Note not found",
		})
	}

	if !note.IsPublic {
		user, err := utils.GetUserByEmailFromContext(c)
		if err != nil || note.UserID != user.ID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  401,
				"message": "Unauthorized - Private note",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   200,
		"message":  "Success",
		"metadata": note,
	})
}
