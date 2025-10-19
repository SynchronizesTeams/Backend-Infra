package main

import (
	"fmt"
	"go-api-infra/database"
	"go-api-infra/models"
	"go-api-infra/routes"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using system env")
	}
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
		&models.GuestBook{},
		&models.Mading{},
		&models.Teacher{},
		&models.Eskul{},
		&models.Achievement{},
		&models.Portal{},
		&models.Industry{},
		&models.Certification{},
		&models.Testimonial{},
	)

	// 3️⃣ Init Fiber
	app := fiber.New()
	allowedOrigins := []string{
		"https://infra-adeli.synchronizeteams.my.id",
		"https://smkpluspelitanusantara.sch.id",
		"http://localhost:3000",
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(allowedOrigins, ","), 
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))


	// 4️⃣ Serve static uploads
	app.Static("/uploads", "./uploads")

	// 5️⃣ Setup API Routes
	routes.SetupRoutes(app)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // fallback kalau .env kosong
	}

	// 6️⃣ Start Server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("🚀 Server running on %s\n", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("❌ Failed to start Fiber server: %v", err)
	}

}
