package routes

import (
	"app/server/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(r fiber.Router) {
	r.Static("/", "../client/dist")

	api := r.Group("/api")
	setupPlayerRoutes(api.Group("/player"))
	api.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("../404.html")
	})
}

func setupPlayerRoutes(r fiber.Router) {
	r.Get("/:id", handlers.GetPlayer())
	r.Post("/", handlers.CreatePlayer())
}
