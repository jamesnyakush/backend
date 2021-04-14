package building

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.Building) (model.Building, error)
	FetchAll() ([]model.Building, error)
	GetByID(uuid.UUID) (model.Building, error)
	AddImage(model.BuildingImage) (model.BuildingImage, error)
	Delete(model.Building) error
	Update(model.Building) error
}

func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}

func (r repository) Add(building model.Building) (model.Building, error) {
	panic("implement me")
}

func (r repository) FetchAll() ([]model.Building, error) {
	panic("implement me")
}

func (r repository) GetByID(uuid uuid.UUID) (model.Building, error) {
	panic("implement me")
}

func (r repository) AddImage(image model.BuildingImage) (model.BuildingImage, error) {
	panic("implement me")
}

func (r repository) Delete(building model.Building) error {
	panic("implement me")
}

func (r repository) Update(building model.Building) error {
	panic("implement me")
}
