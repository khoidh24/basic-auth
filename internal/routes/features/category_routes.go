package features

import (
	category "leanGo/internal/handlers/category"
	"leanGo/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(router fiber.Router) {
	group := router.Group("/category", middleware.ProtectRoutes())

	group.Get("/", category.GetAllCategory)
	group.Get("/:id", category.GetDetailCategory)
	group.Put("/:id", category.UpdateCategory)
	group.Patch("/deactivate", category.DeactiveManyCategory)
	group.Delete("/force", category.HardDeleteManyCategory)
	group.Delete("/force/all", category.HardDeleteAllCategory)
}
