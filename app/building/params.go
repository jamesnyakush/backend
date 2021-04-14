package building

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type BuildingParams struct {
	BuildingName string `json:"building_name" schema:"building_name" form:"building_name"`
	Verified     string `json:"verified" schema:"verified" form:"verified"`
	Description  string `json:"description" schema:"description" form:"description"`
}

func (body BuildingParams) Validate() error {

	err := validation.ValidateStruct(body,
		validation.Field(&body.BuildingName, validation.Required, is.Alphanumeric),
		validation.Field(&body.Verified, validation.Required, is.Int),
		validation.Field(&body.Description, validation.Length(8, 32), is.Alphanumeric),
	)

	return errors.ParseValidationErrorMap(err)
}

type buildingImageRequest struct {
	ImageUrl string `json:"image_url" schema:"image_url" form:"image_url"`
}

func (body *buildingImageRequest) Validate() error {
	err := validation.ValidateStruct(body,
		validation.Field(&body.ImageUrl, validation.Required, is.URL),
	)

	return errors.ParseValidationErrorMap(err)
}
