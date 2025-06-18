package category

import (
	models "leanGo/internal/models/features"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func CreateCategory(c *fiber.Ctx) error {

	user, err := utils.GetUserByEmailFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	var body models.Category
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	body.UserID = user.ID
	body.IsActive = true

	if err := mgm.Coll(&body).Create(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error: Creating category failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Category created successfully",
	})
}
