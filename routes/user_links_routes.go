package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api-infra/handlers"
	jwtware "github.com/gofiber/jwt/v3"
)

func UserLinksRoutes(router fiber.Router) {
	userLinks := router.Group("/user-links")

	userLinks.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	userLinks.Post("/add", handlers.AddLink)
	userLinks.Post("/edit/:id", handlers.EditLinks)
	userLinks.Get("/show/self", handlers.ShowUserLinksSelf)
	userLinks.Delete("/delete/:id", handlers.DeleteLinks)
}
