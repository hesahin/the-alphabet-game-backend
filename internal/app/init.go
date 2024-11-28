package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/dig"
	"the-alphabet-game-backend/api/dataaccess"
	"the-alphabet-game-backend/api/server/controllers"
	"the-alphabet-game-backend/api/services"
	"the-alphabet-game-backend/internal/app/modules"
	"the-alphabet-game-backend/internal/config"
	"the-alphabet-game-backend/pkg"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Name        string
}

func Init(env *config.Env) {
	pkg.SetLogLevel(env)
	fmt.Println(env)

	container := dig.New()

	deps := []Dependency{
		{
			Constructor: func() *config.Env { return env },
			Name:        "Env",
		},
		{
			Constructor: modules.NewQuestionsDb,
			Name:        "QuestionSetDb",
		},
		{
			Constructor: dataaccess.NewQuestionSetDataAccess,
			Interface:   new(dataaccess.IQuestionSetDataAccess),
			Name:        "QuestionSetDataAccess",
		},
		{
			Constructor: controllers.NewQuestionSetController,
			Interface:   new(controllers.IQuestionSetController),
			Name:        "QuestionSetController",
		},
		{
			Constructor: services.NewQuestionSetService,
			Interface:   new(services.IQuestionSetService),
			Name:        "QuestionSetService",
		},
	}

	for _, dep := range deps {
		if dep.Interface == nil {
			//Directly provide the struct without dig.As
			err := container.Provide(dep.Constructor, dig.Name(dep.Name))
			if err != nil {
				log.Panicf("Failed to provide the dependency %s: %v", dep.Name, err)
				panic(err)
			}
		} else {
			//Provide with interface
			err := container.Provide(
				dep.Constructor,
				dig.As(dep.Interface),
				dig.Name(dep.Name),
			)
			if err != nil {
				log.Panicf("Failed to provide the dependency %s: %v", dep.Name, err)
			}
		}
	}

	defer killExternalDependencies(container)

	err := container.Invoke(StartServer)

	if err != nil {
		log.Panicf("Failed to invoke server: %v", err)
		panic(err)
	}

}

func killExternalDependencies(container *dig.Container) {
	container.Invoke(func(questionsDb *modules.QuestionsDb) {
		questionsDb.Stop()
	})
}
