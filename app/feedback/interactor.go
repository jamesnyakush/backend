package feedback

import "github.com/nyumbapoa/backend/app"

type Interactor interface {
}

func NewInteractor() Interactor {
	return &interactor{

	}
}

type interactor struct {
	Config app.Config
}
