package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func CertificationRoutes(router fiber.Router) {
	cert := router.Group("/certification")
	cert.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	cert.Post("/create", handlers.CreateCertification)
	cert.Post("/edit/:id", handlers.EditCertification)
	cert.Get("/show/:id", handlers.ShowCertification)
	cert.Get("/showAll", handlers.ShowAllCertification)
	cert.Delete("/delete/:id", handlers.DeleteCertification)
}