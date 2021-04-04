package utility

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type utilityProviderRequest struct {
	Company     string `json:"company"`
	CountyCode  string `json:"county_code"`
	CountyName  string `json:"county_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (body *utilityProviderRequest) Validate() error {

	err := validation.ValidateStruct(body, validation.Field(&body.Company, validation.Required),
		validation.Field(&body.CountyCode, validation.Required), validation.Field(&body.CountyName, validation.Required),
		validation.Field(&body.PhoneNumber, validation.Required, is.Int),
		validation.Field(&body.Email, validation.Required, is.Email))

	return errors.ParseValidationErrorMap(err)
}

type utilityTypeRequest struct {
	Type string `json:"type"`
}

func (body *utilityTypeRequest) Validate() error {
	err := validation.ValidateStruct(body, validation.Field(&body.Type, validation.Required, is.Alphanumeric))
	return errors.ParseValidationErrorMap(err)
}
