package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func TeacherRoutes(router fiber.Router) {
	teacher := router.Group("/teacher")
	teacher.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	teacher.Post("/create", handlers.CreateTeacher)
	teacher.Post("/edit/:id", handlers.EditTeacher)
	teacher.Delete("/delete/:id", handlers.DeleteTeacher)
	teacher.Get("/showAll", handlers.ShowAllTeachers)
	teacher.Get("/show/:id", handlers.ShowTeacher)
}