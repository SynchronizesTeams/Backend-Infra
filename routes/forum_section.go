package routes

import (
	"go-api-infra/handlers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ForumSectionRoutes(router fiber.Router) {
	section := router.Group("/forum-section")

	section.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	section.Post("/create", handlers.CreateSection)
	section.Post("/edit/:id", handlers.EditSection)
	section.Delete("/delete/:id", handlers.DeleteSection)
	section.Get("/showAll", handlers.ShowAllSection)
}