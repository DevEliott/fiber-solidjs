package routes

import (
	"app/server/api/handlers"
	"app/server/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Static("/", "../client/dist")
	setupWSRoute(app)

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("This is the API")
	})
	api.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	setupPlayerRoutes(api.Group("/player"))

	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusNotFound).SendFile("../404.html")
	// })
}

func setupPlayerRoutes(r fiber.Router) {
	r.Get("/:id", handlers.GetPlayer())
	r.Post("/", handlers.CreatePlayer())
}

func setupWSRoute(app *fiber.App) {
	app.Use("/ws", middleware.IsWebSocketUpgrade)
	app.Get("/ws", handlers.WS())
}
