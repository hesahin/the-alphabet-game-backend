package pkg

import (
	"github.com/gofiber/fiber/v2/log"
	"the-alphabet-game-backend/internal/config"
)

func getLogLevelEnum(env *config.Env) log.Level {
	logLevelString := env.LogLevel
	switch logLevelString {
	case "TRACE":
		return log.LevelTrace
	case "DEBUG":
		return log.LevelDebug
	case "INFO":
		return log.LevelInfo
	case "WARN":
		return log.LevelWarn
	case "ERROR":
		return log.LevelError
	case "FATAL":
		return log.LevelFatal
	case "PANIC":
		return log.LevelPanic
	default:
		return log.LevelError
	}
}

func SetLogLevel(env *config.Env) {
	log.SetLevel(getLogLevelEnum(env))
}
