package routes

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func HealthRoute(group fiber.Router) {
	group.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"time":   time.Now().UnixNano(),
			"status": "OK",
		})
	})
}
