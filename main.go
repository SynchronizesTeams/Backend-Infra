package main

import (
	"go-api-infra/database"
	"go-api-infra/handlers"
	"go-api-infra/models"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	// Auth routes
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Protected routes
	app.Use("/profile", jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))
	app.Get("/profile", handlers.Profile)

	app.Listen(":3000")
}
