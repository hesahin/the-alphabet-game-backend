package main

import (
	"the-alphabet-game-backend/internal/app"
	"the-alphabet-game-backend/internal/config"
)

func main() {
	app.Init(config.LoadEnv())
}
