package house

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app/model"
)

type Interactor interface {
	AddHouse(houseRequest) (model.House, error)
	VerifyHouse(verified bool) (model.House, error)
	FetchAll() ([]model.House, error)
	UpdateHouse(houseId uuid.UUID) error
	DeleteHouse(houseId uuid.UUID) error
}

func NewInteractor() Interactor {
	return &interactor{

	}
}

type interactor struct {
}

func (i interactor) AddHouse(request houseRequest) (model.House, error) {
	panic("implement me")
}

func (i interactor) VerifyHouse(verified bool) (model.House, error) {
	panic("implement me")
}

func (i interactor) FetchAll() ([]model.House, error) {
	panic("implement me")
}

func (i interactor) UpdateHouse(houseId uuid.UUID) error {
	panic("implement me")
}

func (i interactor) DeleteHouse(houseId uuid.UUID) error {
	panic("implement me")
}

