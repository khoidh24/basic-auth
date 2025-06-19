package features

import (
	note "leanGo/internal/handlers/note"
	"leanGo/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterNoteRoutes(router fiber.Router) {
	group := router.Group("/note")

	// Publicly accessible
	group.Get("/:id", middleware.OptionalJWT(), note.GetNoteDetail)

	// Protected routes
	protected := group.Use(middleware.ProtectRoutes())
	protected.Post("/", note.CreateNote)
	protected.Delete("/force", note.HardDeleteManyNotes)
	protected.Put("/:id/status", note.ToggleActiveNote)
	protected.Put("/:id/change-category", note.ChangeNoteCategory)
	protected.Put("/:id", note.UpdateNote)
	protected.Put("/status-notes", note.ToggleActiveNotes)
	protected.Put("/:id/public", note.TogglePublicNote)
}
