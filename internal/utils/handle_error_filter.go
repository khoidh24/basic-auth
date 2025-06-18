package utils

import "github.com/gofiber/fiber/v2"

func HandleFilterError(c *fiber.Ctx, err error, fallbackMsg string) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(fiber.Map{
			"status":  e.Code,
			"message": e.Message,
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  500,
		"message": fallbackMsg,
	})
}
