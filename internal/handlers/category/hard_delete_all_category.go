package category

import (
	features "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func HardDeleteAllCategory(c *fiber.Ctx) error {
	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	// Delete all categories owned by the user
	_, err = mgm.Coll(&features.Category{}).DeleteMany(c.Context(), bson.M{"userId": user.ID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to delete categories",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Deleted all categories successfully",
	})
}
