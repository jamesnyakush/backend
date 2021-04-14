package bill

import "github.com/nyumbapoa/backend/app"

type Interactor interface {
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
