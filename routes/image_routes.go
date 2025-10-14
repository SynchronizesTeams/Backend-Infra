package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api-infra/handlers"
	jwtware "github.com/gofiber/jwt/v3"
)

func ImageRoutes(router fiber.Router) {
	image := router.Group("/image")

	image.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	image.Post("/add", handlers.AddImage)
	image.Delete("/delete/:id", handlers.DeleteImage)
	image.Get("/showAll", handlers.ShowAllImages)
	image.Get("/show", handlers.ShowImagesByCategory)
}
