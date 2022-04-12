package handlers

import (
	"app/server/pkg/player"

	"github.com/gofiber/fiber/v2"
)

func CreatePlayer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody struct{ Name string }
		err := c.BodyParser(&reqBody)
		if err != nil {
			return fiber.ErrBadRequest
		}
		if reqBody.Name == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Must specify param name")
		}
		p := player.CreatePlayer(reqBody.Name)
		return c.Status(fiber.StatusCreated).JSON(p)
	}
}

func GetPlayer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Params("id")
		p, err := player.GetPlayer(ID)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.Status(fiber.StatusOK).JSON(p)
	}
}
