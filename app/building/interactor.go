package building

import (
	"github.com/gofrs/uuid"
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/model"
)

type Interactor interface {
	AddBuilding(BuildingParams) (model.Building, error)
	VerifyBuilding(verified bool) (model.Building, error)
	FetchAll() ([]model.Building, error)
	UpdateBuilding(buildingId uuid.UUID) error
	DeleteBuilding(buildingId uuid.UUID) error
}

func NewInteractor(config app.Config, repo Repository) Interactor {
	return &interactor{
		config:     config,
		repository: repo,
	}
}

type interactor struct {
	config     app.Config
	repository Repository
}

func (i interactor) FetchAll() ([]model.Building, error) {
	panic("implement me")
}

func (i interactor) AddBuilding(params BuildingParams) (model.Building, error) {
	panic("implement me")
}

func (i interactor) VerifyBuilding(verified bool) (model.Building, error) {
	panic("implement me")
}

func (i interactor) UpdateBuilding(buildingId uuid.UUID) error {
	panic("implement me")
}

func (i interactor) DeleteBuilding(buildingId uuid.UUID) error {
	panic("implement me")
}
