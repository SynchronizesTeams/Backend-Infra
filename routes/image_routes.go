package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ImageRoutes(router fiber.Router) {
	image := router.Group("/image")

	image.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	image.Post("/add", handlers.AddImage)
	image.Delete("/delete/:id", handlers.DeleteImage)
	image.Get("/showAll", handlers.ShowAllImages)
	image.Get("/show", handlers.ShowImagesByCategory)
}
