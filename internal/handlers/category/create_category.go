package category

import (
	models "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCategoryRequest struct {
	Name        string `json:"categoryName"`
	Description string `json:"description"`
}

func CreateCategory(c *fiber.Ctx) error {
	// Auth
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Parse body
	var req CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request body",
		})
	}

	// Create Category
	category := &models.Category{
		Name:        req.Name,
		Description: req.Description,
		UserID:      user.ID,
		IsActive:    true,
		NoteIDs:     []primitive.ObjectID{},
		IsPublic:    false,
	}

	if err := mgm.Coll(category).Create(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to create category",
		})
	}

	// Success
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Category created successfully",
		"data":    category,
	})
}
