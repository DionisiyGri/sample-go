package sample

import (
	"github.com/DionisiyGri/sample-go/internal/db/repos"
	"github.com/DionisiyGri/sample-go/internal/model"
	log "github.com/sirupsen/logrus"
)

// App is struct that contains business logic
type App struct {
	some *repos.SomeRepository
}

// NewSampleApp creates and returns sample app
func NewSampleApp(someRepo *repos.SomeRepository) *App {
	return &App{
		some: someRepo,
	}
}

// GetAll contains some logic and execute db layer.
func (a *App) GetAll() ([]model.Some, error) {
	somes, err := a.some.Postgres.GetAll()
	if err != nil {
		log.WithField("method", "GetAll").Errorf("Error getting all  some. err = %s", err)
		return nil, err
	}
	return somes, nil
}

// PostSome contains some logic and execute db layer.
func (a *App) PostSome(item *model.Some) (*model.Some, error) {
	some, err := a.some.Postgres.PostSome(item)
	if err != nil {
		log.WithField("method", "PostSome").Errorf("Error posting some. err = %s", err)
		return nil, err
	}
	return some, nil
}
