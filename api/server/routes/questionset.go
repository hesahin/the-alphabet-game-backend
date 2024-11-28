package routes

import (
	"github.com/gofiber/fiber/v2"
	"the-alphabet-game-backend/api/server/controllers"
)

func QuestionSetRoute(app *fiber.App, controller controllers.IQuestionSetController) {
	app.Get("/questionSet", controller.GetQuestionSet)
}
