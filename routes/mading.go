package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func MadingRoutes(router fiber.Router) {
	mading := router.Group("/mading")
	mading.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	mading.Post("/create", handlers.CreateMading)
	mading.Post("/edit/:id", handlers.EditMading)
	mading.Get("/show/:id", handlers.ShowMading)
	mading.Get("/showAll", handlers.ShowAllMadings)
	// mading.Delete("/delete/:id", handlers.D)
}