package datamodels

type Question struct {
	Id       int    `bson:"id" json:"id"`
	Letter   string `bson:"letter" json:"letter"`
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}
