package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func UserLinksRoutes(router fiber.Router) {
	userLinks := router.Group("/user-links")

	userLinks.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	userLinks.Post("/add", handlers.AddLink)
	userLinks.Post("/edit/:id", handlers.EditLinks)
	userLinks.Get("/show/self", handlers.ShowUserLinksSelf)
	userLinks.Delete("/delete/:id", handlers.DeleteLinks)
}
