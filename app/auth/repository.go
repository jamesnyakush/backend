package auth

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.User) (model.User, error)
	GetByID(uuid.UUID) (model.User, error)
	GetByEmail(string) (model.User, error)
	Update(model.User) error
	Delete(model.User) error
}

func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}

func (r repository) Add(user model.User) (model.User, error) {
	panic("implement me")
}

func (r repository) GetByID(u uuid.UUID) (model.User, error) {
	panic("implement me")
}

func (r repository) GetByEmail(email string) (model.User, error) {
	panic("implement me")
}

func (r repository) Update(user model.User) error {
	panic("implement me")
}

func (r repository) Delete(user model.User) error {
	panic("implement me")
}

