package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func EventRouter(router fiber.Router) {
	event := router.Group("/event")

	event.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	event.Post("/create", handlers.CreateEvent)
	event.Post("/edit/:id", handlers.EditEvent)
	event.Delete("/delete/:id", handlers.DeleteEvent)
	event.Get("/show/:id", handlers.GetEvent)
	event.Get("/showAll", handlers.GetAllEvent)
	event.Put("/change-visibility/:id", handlers.ChangeVisibility)
}


