package routes

import (
	"go-api-infra/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func NewsRoutes(router fiber.Router) {
	news := router.Group("/news")

	news.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	news.Post("/create", handlers.CreateNews)
	news.Post("/update/:id", handlers.UpdateNews)
	news.Get("/getAll", handlers.GetAllNews)
	news.Get("/get/:id", handlers.GetNews)
	news.Delete("/delete/:id", handlers.DeleteNews)
	news.Post("/addView/:id", handlers.AddView)
	news.Post("/changeStatus/:id", handlers.ChangeStatus)
}
