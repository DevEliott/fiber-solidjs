package main

import (
	"app/server/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(), logger.New(), etag.New(), recover.New())

	routes.Setup(app.Group("/api"))

	app.Static("/", "../client/dist")
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("../404.html")
	})

	log.Fatal(app.Listen(":8080"))
}
