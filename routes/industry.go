package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func IndustryRoutes(router fiber.Router)  {
	industry := router.Group("/industry")
	industry.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	industry.Post("/create", handlers.CreateIndustry)
	industry.Post("/edit/:id", handlers.EditIndustry)
	industry.Get("/show/:id", handlers.ShowIndustry)
	industry.Get("/showAll", handlers.ShowAllIndustry)
	industry.Delete("/delete/:id", handlers.DeleteIndustry)
}