package feedback

import (
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.Feedback) (model.Feedback, error)
	FetchAll() ([]model.Feedback, error)
}

func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}

func (r repository) Add(feedback model.Feedback) (model.Feedback, error) {
	panic("implement me")
}

func (r repository) FetchAll() ([]model.Feedback, error) {
	panic("implement me")
}

