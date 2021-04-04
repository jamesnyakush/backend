package feedback

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nyumbapoa/backend/app/errors"
)

type feedbackRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func (body *feedbackRequest) Validate() error {

	err := validation.ValidateStruct(body, validation.Field(&body.Subject, validation.Required),
		validation.Field(&body.Description, validation.Required))

	return errors.ParseValidationErrorMap(err)
}
