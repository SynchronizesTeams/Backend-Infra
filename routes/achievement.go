package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AchievementRoutes(router fiber.Router) {
	achieve := router.Group("/achievement")
	achieve.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	achieve.Post("/create", handlers.CreateAchievement)
	achieve.Post("/edit/:id", handlers.EditAchievement)
	achieve.Get("/show/:id", handlers.ShowAchievement)
	achieve.Get("/showAll", handlers.ShowAllAchievement)
	achieve.Delete("/delete/:id", handlers.DeleteAchievement)
}