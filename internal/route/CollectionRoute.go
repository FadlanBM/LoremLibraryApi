package route

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func CollectionRoute(r fiber.Router) {
	app := r.Group("/collection")
	app.Use(middleware.APIKeyAuthMiddlewareMePeminjam)
	app.Get("/", controller.LoveIndex)
	app.Get("/check/:id", controller.CheckLoveIndex)
	app.Post("/create/:id", controller.CollectionCreate)
	app.Delete("/delete/:id", controller.CollectionDelete)
}
