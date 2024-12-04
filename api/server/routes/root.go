package routes

import (
	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(group fiber.Router) {
	group.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "Uppss, endpoint is not found",
			})
		},
	)
}
