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
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(cors.New(), logger.New(), etag.New(), recover.New())
	routes.Setup(app)
	log.Fatal(app.Listen(":8080"))
}
