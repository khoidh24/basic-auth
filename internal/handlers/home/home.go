package home

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	name := c.Locals("userName").(string)
	return c.JSON(fiber.Map{"message": "Welcome " + name})
}
