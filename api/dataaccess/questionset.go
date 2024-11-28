package dataaccess

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/dig"
	"the-alphabet-game-backend/api/datamodels"
	"the-alphabet-game-backend/internal/app/modules"
)

type IQuestionSetDataAccess interface {
	GetQuestionSet(date string) (*datamodels.QuestionSet, error)
}

type QuestionSetDataAccess struct {
	QuestionsDb *modules.QuestionsDb
}

type QuestionSetDataAccessDependencies struct {
	dig.In

	QuestionsDb *modules.QuestionsDb `name:"QuestionSetDb"`
}

func NewQuestionSetDataAccess(deps QuestionSetDataAccessDependencies) *QuestionSetDataAccess {
	return &QuestionSetDataAccess{
		QuestionsDb: deps.QuestionsDb,
	}
}

func (dataAccess *QuestionSetDataAccess) GetQuestionSet(date string) (*datamodels.QuestionSet, error) {
	collection := dataAccess.QuestionsDb.Collection
	filter := bson.D{{"date", date}}

	var questionSet datamodels.QuestionSet
	queryResult := collection.FindOne(context.Background(), filter)
	err := queryResult.Decode(&questionSet)

	if err != nil {
		log.Errorf("Failed to find document: %v", err)
		return nil, err
	}

	if &questionSet == nil {
		err = fmt.Errorf("no question set found")
		return nil, err
	}

	return &questionSet, nil
}
