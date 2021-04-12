package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type RegisterParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (body RegisterParams) Validate() error {

	err := validation.ValidateStruct(body,
		validation.Field(&body.Name, validation.Length(3, 32), is.ASCII),
		validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.Phone, validation.Required, is.Int),
		validation.Field(&body.Password, validation.Required, validation.Length(8, 32), is.Alphanumeric),
	)

	return errors.ParseValidationErrorMap(err)
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (body LoginParams) Validate() error {
	err := validation.ValidateStruct(body, validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.Password, validation.Length(8, 32), is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)
}
