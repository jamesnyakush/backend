package role

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.Role) (model.Role, error)
	FetchAll() ([]model.Role, error)
	GetById(uuid.UUID) (model.Role, error)
	Update(model.Role) error
	Delete(model.Role) error
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

func (r repository) Add(role model.Role) (model.Role, error) {
	panic("implement me")
}

func (r repository) FetchAll() ([]model.Role, error) {
	panic("implement me")
}

func (r repository) GetById(roleId uuid.UUID) (model.Role, error) {
	panic("implement me")
}

func (r repository) Update(role model.Role) error {
	panic("implement me")
}

func (r repository) Delete(role model.Role) error {
	panic("implement me")
}
