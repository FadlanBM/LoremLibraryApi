package route

import (
	"github.com/gofiber/fiber/v2"
)

func Index(r fiber.Router) {
	app := r.Group("/api")
	AuhtRoute(app)
	BookRoute(app)
	LendingRoute(app)
	CollectionRoute(app)
	CommentRoute(app)
	MetRoute(app)
}
