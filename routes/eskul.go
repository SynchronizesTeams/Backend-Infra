package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func EskulRoutes(router fiber.Router) {
	eskul := router.Group("/eskul")

	eskul.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	eskul.Post("/create", handlers.CreateEskul)
	eskul.Post("/edit/:id", handlers.EditEskul)
	eskul.Get("/show/:id", handlers.ShowEskul)
	eskul.Get("/showAll", handlers.ShowAllEskul)
	eskul.Delete("/delete/:id", handlers.DeleteEskul)
}