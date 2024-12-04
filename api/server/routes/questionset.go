package routes

import (
	"github.com/gofiber/fiber/v2"
	"the-alphabet-game-backend/api/server/controllers"
)

func QuestionSetRoute(group fiber.Router, controller controllers.IQuestionSetController) {
	api := group.Group("/questions")
	api.Get("/questionSet", controller.GetQuestionSet)
}
