package routes

import (
	"go-api-infra/handlers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func PublicRoutes(router fiber.Router) {
	public := router.Group("/public")
	public.Use(cache.New(cache.Config{
		Expiration: 60 * time.Second,
		CacheControl: true,
	}))
	news := public.Group("/news")
	mading := public.Group("/mading")
	image := public.Group("/image")
	teacher := public.Group("/teacher")
	event := public.Group("/event")
	guest := public.Group("/guest")
	eskul := public.Group("/eskul")
	achieve := public.Group("/achievement")
	user := public.Group("/user")
	user_links := public.Group("/user-link")
	testi := public.Group("/testi")
	industry := public.Group("/industry")
	cert := public.Group("/certification")
	portal := public.Group("/portal")

	//Achievement
	achieve.Get("/show/:id", handlers.ShowAchievement)
	achieve.Get("/showAll", handlers.ShowAllAchievement)

	// Certification
	cert.Get("/show/:id", handlers.ShowCertification)
	cert.Get("/showAll", handlers.ShowAllCertification)

	//Eskul
	eskul.Get("/show/:id", handlers.ShowEskul)
	eskul.Get("/showAll", handlers.ShowAllEskul)

	//Event
	event.Get("/show/:id", handlers.GetEvent)
	event.Get("/showAll", handlers.GetAllEvent)
	event.Get("/show/date/:date", handlers.GetEventByDate)

	//Guest
	guest.Get("/showAll", handlers.GetAllGuestBook)
	guest.Get("/show/:id", handlers.GetGuestBook)

	// image
	image.Get("/showAll", handlers.ShowAllImages)
	image.Get("/show", handlers.ShowImagesByCategory)

	// industry
	industry.Get("/show/:id", handlers.ShowIndustry)
	industry.Get("/showAll", handlers.ShowAllIndustry)

	// Mading
	mading.Get("/show/:id", handlers.ShowMading)
	mading.Get("/showAll", handlers.ShowAllMadings)

	// news
	news.Get("/getAll", handlers.GetAllNews)
	news.Get("/get/:id", handlers.GetNews)

	//portal
	portal.Get("/show/:id", handlers.ShowPortal)
	portal.Get("/showAll", handlers.ShowAllPortal)

	// teacher
	teacher.Get("/showAll", handlers.ShowAllTeachers)
	teacher.Get("/show/:id", handlers.ShowTeacher)

	// testimonial
	testi.Get("/show/:id", handlers.ShowTestimonial)
	testi.Get("/showAll", handlers.ShowAllTestimonial)

	//user links
	user_links.Get("/show/:id", handlers.ShowUserLinks)

	//User Profile
	user.Get("profile/:id", handlers.ProfilePublic)
}