package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	AuthRoutes(api)
	ImageRoutes(api)
	NewsRoutes(api)
	UserLinksRoutes(api)
	ForumSectionRoutes(api)
	ForumPostRoutes(api)
	EventRouter(api)
	GuestBookRoutes(api)
}
