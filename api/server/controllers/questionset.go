package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
	"the-alphabet-game-backend/api/schema"
	"the-alphabet-game-backend/api/services"
)

type IQuestionSetController interface {
	GetQuestionSet(c *fiber.Ctx) error
}

type QuestionSetController struct {
	questionSetService services.IQuestionSetService
}

type QuestionSetControllerDependencies struct {
	dig.In

	QuestionSetService services.IQuestionSetService `name:"QuestionSetService"`
}

func NewQuestionSetController(deps QuestionSetControllerDependencies) *QuestionSetController {
	return &QuestionSetController{
		questionSetService: deps.QuestionSetService,
	}
}

func (controller *QuestionSetController) GetQuestionSet(c *fiber.Ctx) error {
	questionSet, err := controller.questionSetService.GetQuestionSet()

	if err != nil {
		return c.JSON(schema.QuestionsResponse{Success: false, Message: err.Error(), Questions: *questionSet})
	}
	return c.JSON(schema.QuestionsResponse{Success: true, Message: "Questions retrieved successfully.", Questions: *questionSet})
}
