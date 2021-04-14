package county

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nyumbapoa/backend/app/errors"
)

type countyRequest struct {
	CountyName string `json:"county_name" schema:"county_name" form:"county_name"`
	CountyCode string `json:"county_code" schema:"county_code" form:"county_code"`
}

func (body countyRequest) Validate() error {

	err := validation.ValidateStruct(body,
		validation.Field(&body.CountyName, validation.Required),
		validation.Field(&body.CountyCode, validation.Required, validation.Length(1, 3)),
	)

	return errors.ParseValidationErrorMap(err)
}
