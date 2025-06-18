package routes

import (
	"leanGo/internal/handlers/auth"
	"leanGo/internal/handlers/docs"
	"leanGo/internal/handlers/home"
	"leanGo/internal/middleware"
	"leanGo/internal/routes/features"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	// Auth
	v1.Post("/signup", auth.SignUp)
	v1.Post("/login", auth.Login)

	// Home
	v1.Get("/home", middleware.ProtectRoutes(), home.Home)

	// Category Routes
	features.RegisterCategoryRoutes(v1)

	// Swagger Docs
	v1.Get("/swagger.json", docs.SwaggerJSON)
	v1.Get("/reference", docs.ReferencePage)
}
