package route

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuhtRoute(r fiber.Router) {
	app := r.Group("/auth")

	app.Post("/", controller.AuthController)
	app.Post("/register", controller.RegisterController)
	app.Put("/register/image/:id", controller.RegisterAvatar)
	appme := r.Group("/auth")
	appme.Use(middleware.APIKeyAuthMiddlewareMePeminjam)
	appme.Get("/me", controller.MePeminjamController)
}
