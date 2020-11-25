package house

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
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

type houseImageRequest struct {
	HouseId  string `json:"house_id"`
	ImageUrl string `json:"image_url"`
}

type houseTypeRequest struct {
	Type string `json:"type"`
}

func (body *houseRequest) Bind(r *http.Request) error {
	body.UserId = strings.TrimSpace(body.UserId)
	body.BuildingId = strings.TrimSpace(body.BuildingId)
	body.HouseTypeId = strings.TrimSpace(body.HouseTypeId)
	body.HouseNumber = strings.TrimSpace(body.HouseNumber)
	body.Title = strings.TrimSpace(body.Title)
	body.Room = strings.TrimSpace(body.Room)
	body.Verified = strings.TrimSpace(body.Verified)
	body.Occupied = strings.TrimSpace(body.Occupied)
	body.Description = strings.TrimSpace(body.Description)

	return validation.ValidateStruct(body, validation.Field(&body.UserId, validation.Required, is.UUID),
		validation.Field(&body.BuildingId, validation.Required, is.UUID), validation.Field(&body.HouseTypeId, validation.Required, is.UUID),
		validation.Field(&body.HouseNumber, validation.Required, is.Int), validation.Field(&body.Title, validation.Required, is.Alphanumeric),
		validation.Field(&body.Room, validation.Required, is.Alphanumeric), validation.Field(&body.Verified, is.Int),
		validation.Field(&body.Occupied, validation.Required, is.Int), validation.Field(&body.Description, validation.Length(8, 32), is.Alphanumeric))
}

func (body *houseImageRequest) Bind(r *http.Request) error {
	body.HouseId = strings.TrimSpace(body.HouseId)
	body.ImageUrl = strings.TrimSpace(body.ImageUrl)

	return validation.ValidateStruct(body, validation.Field(&body.HouseId, validation.Required, is.UUID),
		validation.Field(&body.ImageUrl, validation.Required, is.URL))
}

func (body *houseTypeRequest) Bind(r *http.Request) error {
	body.Type = strings.TrimSpace(body.Type)
	return validation.ValidateStruct(body, validation.Field(&body.Type, validation.Required, is.Alphanumeric))
}

func (rs Resource) HandleGetHouses(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleGetHouse(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleGetHouseImage(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleGetHouseType(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleCreateHouse(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleCreateHouseImage(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleCreateHouseType(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleUpdateHouse(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleUpdateHouseType(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleUpdateHouseImage(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleDeleteHouse(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleDeleteHouseImage(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleDeleteHouseType(w http.ResponseWriter, r *http.Request) {
}
