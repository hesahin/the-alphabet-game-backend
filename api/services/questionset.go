package services

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/dig"
	"the-alphabet-game-backend/api/dataaccess"
	"the-alphabet-game-backend/api/datamodels"
	"time"
)

type IQuestionSetService interface {
	GetQuestionSet() (*datamodels.QuestionSet, error)
}

type QuestionSetService struct {
	questionSetDataAccess dataaccess.IQuestionSetDataAccess
}

type QuestionSetServiceDependencies struct {
	dig.In
	QuestionSetDataAccess dataaccess.IQuestionSetDataAccess `name:"QuestionSetDataAccess"`
}

func NewQuestionSetService(deps QuestionSetServiceDependencies) *QuestionSetService {
	return &QuestionSetService{
		questionSetDataAccess: deps.QuestionSetDataAccess,
	}
}

func (service *QuestionSetService) GetQuestionSet() (*datamodels.QuestionSet, error) {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	log.Info(date)
	questionSet, err := service.questionSetDataAccess.GetQuestionSet(date)

	if err != nil {
		return nil, err
	}

	return questionSet, nil
}
