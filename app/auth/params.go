package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type RegisterParams struct {
	Email    string `json:"email" schema:"email" form:"email"`
	Phone    string `json:"phone" schema:"phone" form:"phone"`
	Password string `json:"password" schema:"password" form:"password"`
}

func (body RegisterParams) Validate() error {

	err := validation.ValidateStruct(body,
		validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.Phone, validation.Required, is.Int),
		validation.Field(&body.Password, validation.Required, validation.Length(8, 32), is.Alphanumeric),
	)

	return errors.ParseValidationErrorMap(err)
}

type LoginParams struct {
	Email    string `json:"email" schema:"email" form:"email"`
	Password string `json:"password" schema:"password" form:"password"`
}

func (body LoginParams) Validate() error {
	err := validation.ValidateStruct(body, validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.Password, validation.Length(8, 32), is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)
}
