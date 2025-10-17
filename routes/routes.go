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
	MadingRoutes(api)
	TeacherRoutes(api)
	EskulRoutes(api)
	AchievementRoutes(api)
	PortalRoutes(api)
	IndustryRoutes(api)
	CertificationRoutes(api)
	TestimonialRoutes(api)
	PublicRoutes(api)

	app.Get("/ping", func(c *fiber.Ctx) error {
	return c.SendString("pong")
	})

}
