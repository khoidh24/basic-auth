package main

import (
	"fmt"
	"log"

	configs "leanGo/config"
	"leanGo/internal/database"
	"leanGo/internal/routes"
	"leanGo/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load config
	configs.LoadConfig()

	// Connect DB
	database.ConnectMongo()

	// Routes
	routes.Routes(app)

	// Start server with dynamic port handling
	startPort := utils.GetInitialPort(configs.Port)
	finalPort := utils.FindAvailablePort(startPort)

	log.Printf("Server starting on http://localhost:%d", finalPort)
	if err := app.Listen(fmt.Sprintf(":%d", finalPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
