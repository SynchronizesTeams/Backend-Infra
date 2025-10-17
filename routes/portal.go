package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func PortalRoutes(router fiber.Router) {
	portal := router.Group("/portal")
	portal.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	portal.Post("/create", handlers.CreatePortal)
	portal.Post("/edit/:id", handlers.EditPortal)
	portal.Get("/show/:id", handlers.ShowPortal)
	portal.Get("/showAll", handlers.ShowAllPortal)
	portal.Delete("/delete/:id", handlers.DeletePortal)
}