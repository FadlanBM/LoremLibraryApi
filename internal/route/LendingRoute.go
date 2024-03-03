package route

import (
	"github.com/api/internal/controller"
	"github.com/api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func LendingRoute(r fiber.Router) {
	app := r.Group("/lending")

	app.Use(middleware.APIKeyAuthMiddlewareMePeminjam)
	app.Get("/", controller.LendingIndex)
	app.Post("/create", controller.LendingsCreate)
	app.Post("/list", controller.ListLendingCreate)
	app.Get("/:id", controller.DetailLending)
	app.Get("/history/:id", controller.HistroyLending)
	app.Delete("/delete/:id", controller.DeleteLending)
}
