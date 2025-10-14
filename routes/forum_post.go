package routes

import (
	"go-api-infra/handlers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ForumPostRoutes(router fiber.Router) {
	post := router.Group("/forum-post")
	post.Use(jwtware.New(jwtware.Config{
		SigningKey: handlers.SecretKey,
	}))

	post.Post("/create", handlers.CreatePost)
	post.Post("/edit/:id", handlers.EditPost)
	post.Delete("/delete/:id", handlers.DeletePost)
	post.Get("/showAll", handlers.GetAllPost)
	post.Get("/show/:id", handlers.ShowPostById)
	post.Put("/changeStatus", handlers.ChangePostStatus)
	post.Put("/upvote/:id", handlers.UpvotePost)
	post.Put("/downvote/:id", handlers.DownvotePost)
	post.Post("/reply/:id", handlers.CreateReply)
	post.Get("/post-replies/:id", handlers.GetPostWithReplies)
	post.Get("/replies/:id", handlers.GetNestedReplies)
}