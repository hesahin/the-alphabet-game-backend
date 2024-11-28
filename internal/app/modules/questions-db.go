package modules

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
	"the-alphabet-game-backend/internal/config"
)

type QuestionsDb struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

type QuestionDbDependencies struct {
	dig.In

	Env *config.Env `name:"Env"`
}

func NewQuestionsDb(deps QuestionDbDependencies) (*QuestionsDb, error) {
	credential := options.Credential{Username: "root", Password: "password"}
	clientOptions := options.Client().ApplyURI(deps.Env.QuestionsDbConfig.MongoQuestionsDbUrl).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	db := client.Database(deps.Env.QuestionsDbConfig.MongoQuestionsDbName)
	collection := db.Collection(deps.Env.QuestionsDbConfig.MongoQuestionsDbCollectionName)

	res := &QuestionsDb{
		Client:     client,
		Database:   db,
		Collection: collection,
	}
	return res, nil
}

func (mongoDb *QuestionsDb) Stop() {
	if mongoDb.Client != nil {
		if err := mongoDb.Client.Disconnect(context.TODO()); err != nil {
			log.Errorf("Error while disconnecting from QuestionsDb: %v", err)
		} else {
			log.Info("Disconnected from QuestionsDb")
		}
	} else {
		log.Warn("Already disconnected from QuestionsDb")
	}
}
