package route

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommentRoute(r fiber.Router) {
	app := r.Group("/comment")

	app.Use(middleware.APIKeyAuthMiddlewareMePeminjam)
	app.Post("/", controller.CommentCreate)
	app.Get("/:id", controller.CommentIndex)
	app.Get("/check/:id", controller.CommentCheck)
}
