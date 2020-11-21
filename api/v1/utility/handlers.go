package utility

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"strings"
)

type utilityProviderRequest struct {
	Company     string `json:"company"`
	CountyCode  string `json:"county_code"`
	CountyName  string `json:"county_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type utilityTypeRequest struct {
	Type string `json:"type"`
}

func (body *utilityProviderRequest) Bind(r http.Request) error {
	body.Company = strings.TrimSpace(body.Company)
	body.CountyCode = strings.TrimSpace(body.CountyCode)
	body.CountyName = strings.TrimSpace(body.CountyName)
	body.PhoneNumber = strings.TrimSpace(body.PhoneNumber)
	body.Email = strings.TrimSpace(body.Email)

	return validation.ValidateStruct(body, validation.Field(&body.Company, validation.Required),
		validation.Field(&body.CountyCode, validation.Required), validation.Field(&body.CountyName, validation.Required),
		validation.Field(&body.PhoneNumber, validation.Required, is.Int),
		validation.Field(&body.Email, validation.Required, is.Email))
}

func (body *utilityTypeRequest) Bind(r http.Request) error {
	body.Type = strings.TrimSpace(body.Type)

	return validation.ValidateStruct(body, validation.Field(&body.Type, validation.Required, is.Alphanumeric))
}
