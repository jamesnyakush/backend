package house

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.House) (model.House, error)
	FetchAll() ([]model.House, error)
	GetById(uuid.UUID)
	Update(model.House) error
	Delete(model.House) error
}

// NewRepository creates and returns a new instance of admin repository
func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}

func (r repository) Add(house model.House) (model.House, error) {
	return house, nil
}

func (r repository) FetchAll() ([]model.House, error) {
	panic("implement me")
}

func (r repository) GetById(u uuid.UUID) {
	panic("implement me")
}

func (r repository) Update(house model.House) error {
	panic("implement me")
}

func (r repository) Delete(house model.House) error {
	panic("implement me")
}
