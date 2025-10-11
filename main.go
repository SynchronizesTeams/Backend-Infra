package main

import (
	// "time"
	"go-api-infra/database"
	"go-api-infra/handlers"
	"go-api-infra/models"
	//  "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2/middleware/cache"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(
		&models.User{},
		&models.News{},
		&models.Image{},
		&models.UserLinks{},
	)

	app := fiber.New()
	app.Static("/uploads", "./uploads")

	// Route api group
	api := app.Group("/api/v1")

	//Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	//Protected Route
	auth.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))
	auth.Get("/profile", handlers.Profile)

	//Route Group /images
	image := api.Group("/image")
	image.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	image.Post("/add", handlers.AddImage)
	image.Delete("/delete/:id", handlers.DeleteImage)
	image.Get("/showAll", handlers.ShowAllImages)
	image.Get("show", handlers.ShowImagesByCategory)

	// Route group /news
	news := api.Group("/news")
	news.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	news.Post("/create", handlers.CreateNews)
	news.Post("/update/:id", handlers.UpdateNews)
	news.Get("/getAll", handlers.GetAllNews)
	news.Get("/get/:id", handlers.GetNews)
	news.Delete("/delete/:id", handlers.DeleteNews)
	news.Post("/addView/:id", handlers.AddView)
	news.Post("/changeStatus/:id", handlers.ChangeStatus)

	// Route group user links
	userLinks := api.Group("/user-links")
	userLinks.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	userLinks.Post("/add", handlers.AddLink)

	app.Listen(":3000")
}
