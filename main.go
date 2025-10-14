package main

import (
	"go-api-infra/database"
	"go-api-infra/models"
	"go-api-infra/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1️⃣ Connect Database
	database.Connect()

	// 2️⃣ Auto migrate all models
	database.DB.AutoMigrate(
		&models.User{},
		&models.News{},
		&models.Image{},
		&models.UserLinks{},
		&models.ForumPost{},
		&models.ForumReply{},
		&models.ForumSection{},
		&models.Event{},
	)

	// 3️⃣ Init Fiber
	app := fiber.New()

	// 4️⃣ Serve static uploads
	app.Static("/uploads", "./uploads")

	// 5️⃣ Setup API Routes
	routes.SetupRoutes(app)

	// 6️⃣ Start Server
	app.Listen(":3000")
}
