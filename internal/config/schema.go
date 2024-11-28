package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv            string            `mapstructure:"APP_ENV"`
	Port              int               `mapstructure:"PORT"`
	QuestionsDbConfig QuestionsDbConfig `mapstructure:",squash"`
	LogLevel          string            `mapstructure:"LOG_LEVEL"`
}

type QuestionsDbConfig struct {
	MongoQuestionsDbUrl            string `mapstructure:"MONGO_QUESTIONSDB_URL"`
	MongoQuestionsDbName           string `mapstructure:"MONGO_QUESTIONSDB_NAME"`
	MongoQuestionsDbCollectionName string `mapstructure:"MONGO_QUESTIONSDB_COLLECTION"`
}

func LoadEnv() *Env {
	env := Env{}

	viper.SetConfigFile("./.env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't	 find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Info("The App is running in development env")
	}

	return &env
}
