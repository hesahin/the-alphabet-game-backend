package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"syscall"
	"the-alphabet-game-backend/api/middlewares"
	"the-alphabet-game-backend/api/server/controllers"
	"the-alphabet-game-backend/api/server/routes"
	"the-alphabet-game-backend/internal/config"
)

type Server struct {
	app                   *fiber.App
	env                   *config.Env
	questionSetController controllers.IQuestionSetController
}

type ServerDependencies struct {
	dig.In

	Env                   *config.Env                        `name:"Env"`
	QuestionSetController controllers.IQuestionSetController `name:"QuestionSetController"`
}

func StartServer(deps ServerDependencies) {
	server := &Server{
		app:                   fiber.New(),
		env:                   deps.Env,
		questionSetController: deps.QuestionSetController,
	}

	server.bindMiddlewares()

	server.bindRoutes()

	server.gracefulShutdown()

	err := server.app.Listen((fmt.Sprintf(":%d", server.env.Port)))
	if err != nil {
		log.Panicf("Failed to start the server: %s", err)
		panic(err)
	}
}

func (server *Server) bindRoutes() {
	routes.HealthRoute(server.app)
	routes.QuestionSetRoute(server.app, server.questionSetController)
	routes.NotFoundRoute(server.app)
}

func (server *Server) gracefulShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		log.Info("Shutting down the server...")
		_ = server.app.Shutdown()
	}()
}

func (server *Server) bindMiddlewares() {
	server.app.Use(middlewares.NewLogger())
	server.app.Use(helmet.New())
	server.app.Use(recover.New())
	server.app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Specify your Next.js app's origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))
	server.app.Use(requestid.New())
}
