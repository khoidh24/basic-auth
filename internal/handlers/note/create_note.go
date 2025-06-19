package note

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNote handles creation of a new note by an authenticated user.
func CreateNote(c *fiber.Ctx) error {
	// Get user from context
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	type Body struct {
		Title      string `json:"noteTitle"`
		Content    string `json:"content"`
		CoverImage string `json:"coverImage"`
		CategoryID string `json:"categoryId"`
	}

	var body Body
	if reqErr := c.BodyParser(&body); reqErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	// Validate required fields
	if body.Title == "" || body.CategoryID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Title and categoryId are required",
		})
	}

	// Convert CategoryID to ObjectID
	categoryObjID, err := primitive.ObjectIDFromHex(body.CategoryID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid category ID",
		})
	}

	// Find the category
	category := &features.Category{}
	if err := mgm.Coll(category).FindByID(categoryObjID, category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to find category",
		})
	}

	// Create new note instance
	newNote := &features.Note{
		Title:      body.Title,
		Content:    body.Content,
		CoverImage: body.CoverImage,
		UserID:     user.ID,
		CategoryID: categoryObjID,
		IsPublic:   false,
		IsActive:   true,
	}

	// Save to DB
	if err := mgm.Coll(newNote).Create(newNote); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// Append newNote.ID to category.NoteIDs if not exists
	category.NoteIDs = append(category.NoteIDs, newNote.ID)

	// Update the category
	if err := mgm.Coll(category).Update(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to update category",
		})
	}

	// Success
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "Note created successfully",
		"data":    newNote,
	})
}
