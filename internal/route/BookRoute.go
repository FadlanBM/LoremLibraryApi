package route

import (
	"github.com/api/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func BookRoute(r fiber.Router) {
	app := r.Group("/book")

	app.Get("/", controller.BookIndex)
}
