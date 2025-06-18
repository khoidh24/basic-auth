package category

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/services"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDetailCategory(c *fiber.Ctx) error {
	// Parse ID
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid ID",
		})
	}

	// Find category by ID
	category := &features.Category{}
	if notFoundCategoryErr := mgm.Coll(category).FindByID(objID, category); notFoundCategoryErr != nil || !category.IsActive {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": "Category not found",
		})
	}

	// Get user
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Check ownership
	if category.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  403,
			"message": "Forbidden",
		})
	}

	// Build filters from query
	result, err := services.FilterBuilder(c, services.FilterOptions{
		DefaultLimit: 10,
		AllowSortBy:  []string{"createdAt", "name"},
		ExtraFilters: map[string]string{},
	})
	if err != nil {
		return utils.HandleFilterError(c, err, "Failed to parse filters")
	}

	// Filter by user + category
	result.Filter["userId"] = user.ID
	result.Filter["categoryId"] = category.ID

	// Count notes
	total, err := mgm.Coll(&features.Note{}).CountDocuments(c.Context(), result.Filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Error counting notes",
		})
	}

	// Find notes
	cursor, err := mgm.Coll(&features.Note{}).Find(c.Context(), result.Filter, result.FindOpts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Error fetching notes",
		})
	}
	defer cursor.Close(c.Context())

	var notes []features.Note
	if err := cursor.All(c.Context(), &notes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Error decoding notes",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"metadata": fiber.Map{
			"data":  notes,
			"page":  result.Pagination.Page,
			"limit": result.Pagination.Limit,
			"total": total,
		},
	})
}
