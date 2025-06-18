// /routes/features/category_routes.go
package features

import (
	category "leanGo/internal/handlers/category"
	"leanGo/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(router fiber.Router) {
	group := router.Group("/category")

	// Public route (public detail without auth)
	group.Get("/:id", category.GetDetailCategory)

	// Protected routes
	protected := group.Use(middleware.ProtectRoutes())

	protected.Get("/", category.GetAllCategory)
	protected.Post("/", category.CreateCategory)
	protected.Put("/:id", category.UpdateCategoryInfo)
	protected.Put("/deactivate", category.DeactiveManyCategory)
	protected.Delete("/force", category.HardDeleteManyCategory)
	protected.Delete("/force/all", category.HardDeleteAllCategory)
	protected.Put("/:id/status", category.ToggleActiveCategory)
	protected.Put("/:id/public", category.TogglePublicCategory)
}
