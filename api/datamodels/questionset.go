package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuestionSet struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	ValidDate string             `bson:"date" json:"date"`
	Questions []Question         `bson:"questions" json:"questions"`
}
