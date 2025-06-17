package routes

import (
	configs "leanGo/config"
	"leanGo/internal/handlers/auth"
	"leanGo/internal/handlers/home"
	"leanGo/internal/middleware"
	"log"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	v1 := app.Group("/api/v1")

	v1.Post("/signup", auth.SignUp)
	v1.Post("/login", auth.Login)
	v1.Get("/home", middleware.ProtectRoutes(), home.Home)

	v1.Get("/swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	v1.Get("/swagger", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "http://localhost:" + configs.Port + "/api/v1/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Habify API",
			},
			DarkMode: true,
		})

		if err != nil {
			log.Printf("❌ Scalar error: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate Scalar UI")
		}

		return c.Type("html").SendString(htmlContent)
	})
}
