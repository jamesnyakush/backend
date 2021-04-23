package auth

import (
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/errors"
	"github.com/nyumbapoa/backend/app/helpers"
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

	// hash admin password before adding to db.
	passwordHash, err := helpers.HashPassword(user.Password)

	// if we get an error, it means our hashing func dint work
	if err != nil {
		return model.User{}, errors.Error{Err: err, Code: errors.EINTERNAL}
	}

	// change password to hashed string
	user.Password = passwordHash
	usr, err := i.repository.Add(user)

	if err != nil {
		return model.User{}, err
	}

	return usr, nil
}
