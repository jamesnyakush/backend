package building

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"strings"
)

type buildingRequest struct {
	UserId       string `json:"user_id"`
	BuildingName string `json:"building_name"`
	Verified     string `json:"verified"`
	Description  string `json:"description"`
}

type buildingImageRequest struct {
	BuildingId string `json:"building_id"`
	ImageUrl   string `json:"image_url"`
}

type buildingTypeRequest struct {
	Type string `json:"type"`
}

func (body *buildingRequest) Bind(r *http.Request) error {
	body.UserId = strings.TrimSpace(body.UserId)
	body.BuildingName = strings.TrimSpace(body.BuildingName)
	body.Verified = strings.TrimSpace(body.Verified)
	body.Description = strings.TrimSpace(body.Description)

	return validation.ValidateStruct(body, validation.Field(&body.UserId, validation.Required, is.UUID),
		validation.Field(&body.BuildingName, validation.Required, is.Alphanumeric),
		validation.Field(&body.Verified, validation.Required, is.Int), validation.Field(&body.Description, validation.Length(8, 32), is.Alphanumeric))
}

func (body *buildingImageRequest) Bind(r *http.Request) error {
	body.BuildingId = strings.TrimSpace(body.BuildingId)
	body.ImageUrl = strings.TrimSpace(body.ImageUrl)
	return validation.ValidateStruct(body, validation.Field(&body.BuildingId, validation.Required, is.UUID), validation.Field(&body.ImageUrl, validation.Required, is.URL))
}

func (body *buildingTypeRequest) Bind(r *http.Request) error {
	body.Type = strings.TrimSpace(body.Type)
	return validation.ValidateStruct(body, validation.Field(&body.Type, validation.Required, is.Alphanumeric))
}

func (rs Resource) HandleCreateBuilding(w http.ResponseWriter,r *http.Request)  {

	body := buildingRequest{}


	fmt.Println(body)

}

func (rs Resource) HandleUpdateBuiilding(w http.ResponseWriter,r *http.Request)  {

}