package route

import (
	"github.com/api/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func BookRoute(r fiber.Router) {
	app := r.Group("/book")

	app.Get("/", controller.BookIndex)
	app.Get("/image/:path", controller.BookGetImage)
	app.Get("/:id", controller.BookDetail)
	app.Get("/category/:id", controller.CategoryDetail)

}
