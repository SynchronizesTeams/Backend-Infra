package routes

import (
	"go-api-infra/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	auth.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))
	auth.Get("/profile", handlers.Profile)
}
