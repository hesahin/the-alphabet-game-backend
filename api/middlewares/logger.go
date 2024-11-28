package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"the-alphabet-game-backend/internal/config"
)

func NewLogger() fiber.Handler {
	return logger.New(logger.Config{Format: config.LoggerFormat})
}
