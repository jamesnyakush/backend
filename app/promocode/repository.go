package promocode

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.PromoCode) (model.PromoCode, error)
	FetchAll() ([]model.PromoCode, error)
	GetById(uuid.UUID) (model.PromoCode, error)
	Update(model.PromoCode) error
	Delete(model.PromoCode) error
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

func (r repository) Add(code model.PromoCode) (model.PromoCode, error) {
	panic("implement me")
}

func (r repository) FetchAll() ([]model.PromoCode, error) {
	panic("implement me")
}

func (r repository) GetById(u uuid.UUID) (model.PromoCode, error) {
	panic("implement me")
}

func (r repository) Update(code model.PromoCode) error {
	panic("implement me")
}

func (r repository) Delete(code model.PromoCode) error {
	panic("implement me")
}

