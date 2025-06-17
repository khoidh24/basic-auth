package main

import (
	configs "leanGo/config"
	"leanGo/internal/database"
	"leanGo/internal/routes"
	"log"

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

	app.Listen(":" + configs.Port)
	log.Printf("Server listening on http://localhost:%s", configs.Port)
}
