package house

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
	"strings"
)

type houseRequest struct {
	UserId      string `json:"user_id"`
	BuildingId  string `json:"building_id"`
	HouseTypeId string `json:"house_type_id"`
	HouseNumber string `json:"house_number"`
	Title       string `json:"title"`
	Room        string `json:"room"`
	Verified    string `json:"verified"`
	Occupied    string `json:"occupied"`
	Description string `json:"description"`
}

func (body *houseRequest) Validate() error {

	err := validation.ValidateStruct(body, validation.Field(&body.UserId, validation.Required, is.UUID),
		validation.Field(&body.BuildingId, validation.Required, is.UUID), validation.Field(&body.HouseTypeId,
			validation.Required, is.UUID), validation.Field(&body.HouseNumber, validation.Required, is.Int),
		validation.Field(&body.Title, validation.Required, is.Alphanumeric),
		validation.Field(&body.Room, validation.Required, is.Alphanumeric), validation.Field(&body.Verified, is.Int),
		validation.Field(&body.Occupied, validation.Required, is.Int), validation.Field(&body.Description,
			validation.Length(8, 32), is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)
}

type houseImageRequest struct {
	HouseId  string `json:"house_id"`
	ImageUrl string `json:"image_url"`
}

func (body *houseImageRequest) Validate() error {
	body.HouseId = strings.TrimSpace(body.HouseId)
	body.ImageUrl = strings.TrimSpace(body.ImageUrl)

	err := validation.ValidateStruct(body, validation.Field(&body.HouseId, validation.Required, is.UUID),
		validation.Field(&body.ImageUrl, validation.Required, is.URL))

	return errors.ParseValidationErrorMap(err)

}

type houseTypeRequest struct {
	Type string `json:"type"`
}

func (body *houseTypeRequest) Validate() error {
	err := validation.ValidateStruct(body, validation.Field(&body.Type, validation.Required, is.Alphanumeric))

	return errors.ParseValidationErrorMap(err)

}
