package route

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MetRoute(r fiber.Router) {
	app := r.Group("/me")
	app.Use(middleware.APIKeyAuthMiddlewareMePeminjam)
	app.Get("/", controller.MePeminjamController)
	app.Delete("/delete", controller.MeDeleteController)
}
