package category

import (
	auth "leanGo/internal/models/auth"
	features "leanGo/internal/models/features"
	"leanGo/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCategory(c *fiber.Ctx) error {
	// Get user from token
	email := c.Locals("email").(string)
	user := &auth.User{}
	if err := mgm.Coll(user).First(bson.M{"email": email}, user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// FilterBuilder from services
	result, err := services.FilterBuilder(c, services.FilterOptions{
		DefaultLimit: 10,
		AllowSortBy:  []string{"createdAt", "categoryName"},
		ExtraFilters: map[string]string{}, // can be extended later
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to parse filters",
		})
	}

	// Add filter by user
	result.Filter["userId"] = user.ID
	result.Filter["isActive"] = true

	// Total documents
	total, err := mgm.Coll(&features.Category{}).CountDocuments(c.Context(), result.Filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Error counting documents",
		})
	}

	// Find with pagination
	cursor, err := mgm.Coll(&features.Category{}).Find(c.Context(), result.Filter, result.FindOpts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}
	defer cursor.Close(c.Context())

	var categories []features.Category
	if err := cursor.All(c.Context(), &categories); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Error reading data",
		})
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"metadata": fiber.Map{
			"data":  categories,
			"page":  result.Pagination.Page,
			"limit": result.Pagination.Limit,
			"total": total,
		},
	})
}
