package routes

import (
	"app/server/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(r fiber.Router) {
	setupPlayerRoutes(r.Group("/player"))
}

func setupPlayerRoutes(r fiber.Router) {
	r.Get("/:id", handlers.GetPlayer())
	r.Post("/", handlers.CreatePlayer())
}
