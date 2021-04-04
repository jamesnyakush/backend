package building

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type BuildingParams struct {
	UserId       string `json:"user_id"`
	BuildingName string `json:"building_name"`
	Verified     string `json:"verified"`
	Description  string `json:"description"`
}

func (body BuildingParams) Validate() error {

	err := validation.ValidateStruct(body, validation.Field(&body.UserId, validation.Required, is.UUID),
		validation.Field(&body.BuildingName, validation.Required, is.Alphanumeric),
		validation.Field(&body.Verified, validation.Required, is.Int),
		validation.Field(&body.Description, validation.Length(8, 32), is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)
}

type buildingImageRequest struct {
	BuildingId string `json:"building_id"`
	ImageUrl   string `json:"image_url"`
}

func (body *buildingImageRequest) Validate() error {
	err := validation.ValidateStruct(body, validation.Field(&body.BuildingId, validation.Required, is.UUID),
		validation.Field(&body.ImageUrl, validation.Required, is.URL))

	return errors.ParseValidationErrorMap(err)
}

type buildingTypeRequest struct {
	Type string `json:"type"`
}

func (body buildingTypeRequest) Validate() error {
	err := validation.ValidateStruct(body, validation.Field(&body.Type, validation.Required, is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)
}
