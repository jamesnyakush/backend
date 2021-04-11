package auth

import (
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/model"
)

type Interactor interface {
	AuthenticateByEmail(email, password string) (model.User, error)
	Register(RegisterParams) (model.User, error)
}

func NewInteractor(config app.Config, repository Repository) Interactor {
	return &interactor{
		config:     config,
		repository: repository,
	}
}

type interactor struct {
	config     app.Config
	repository Repository
}

func (i interactor) AuthenticateByEmail(email, password string) (model.User, error) {
	panic("implement me")
}

// Register takes in a admin object and adds the admin to db.
func (i interactor) Register(params RegisterParams) (model.User, error) {

	user := model.User{
		Email:    params.Email,
		Password: params.Password,
		Phone:    params.Phone,
	}

	return user, nil
}
