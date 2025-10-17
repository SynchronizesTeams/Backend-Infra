package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func TestimonialRoutes(router fiber.Router) {
	testi := router.Group("/testimonial")
	testi.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	testi.Post("/create", handlers.CreateTestimonial)
	testi.Post("/edit/:id", handlers.EditTestimonial)
	testi.Get("/show/:id", handlers.ShowTestimonial)
	testi.Get("/showAll", handlers.ShowAllTestimonial)
	testi.Delete("/delete/:id", handlers.DeleteTestimonial)
}