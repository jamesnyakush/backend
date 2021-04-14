package permission

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.Permission) (model.Permission, error)
	FetchAll() ([]model.Permission, error)
	GetById(uuid.UUID) (model.Permission, error)
	Update(model.Permission) error
	Delete(model.Permission) error
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

func (r repository) Add(permission model.Permission) (model.Permission, error) {
	panic("implement me")
}

func (r repository) GetById(u uuid.UUID) (model.Permission, error) {
	panic("implement me")
}

func (r repository) FetchAll() ([]model.Permission, error) {
	panic("implement me")
}

func (r repository) Update(permission model.Permission) error {
	panic("implement me")
}

func (r repository) Delete(permission model.Permission) error {
	panic("implement me")
}

