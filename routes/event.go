package routes

import (
	"go-api-infra/handlers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func EventRouter(router fiber.Router) {
	event := router.Group("/event")

	event.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	event.Post("/create", handlers.CreateEvent)
}