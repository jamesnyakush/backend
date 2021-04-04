package role

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type roleRequest struct {
	RoleName string `json:"role_name"`
}

func (body *roleRequest) Validate() error {

	err := validation.ValidateStruct(body, validation.Field(&body.RoleName, validation.Required, is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)
}
