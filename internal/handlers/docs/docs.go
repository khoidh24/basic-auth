package docs

import (
	configs "leanGo/config"
	"log"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
)

func SwaggerJSON(c *fiber.Ctx) error {
	return c.SendFile("./docs/swagger.json")
}

func ReferencePage(c *fiber.Ctx) error {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: configs.Domain + ":" + configs.Port + "/api/v1/swagger.json",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Habify API",
		},
		DarkMode: true,
	})

	if err != nil {
		log.Printf("❌ Scalar error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to generate Scalar UI",
		})
	}

	return c.Type("html").SendString(htmlContent)
}
