package county

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"strings"
)

type countyRequest struct {
	CountyName string `json:"county_name"`
	CountyCode string `json:"county_code"`
}

func (body *countyRequest) Bind(r *http.Request) error {
	body.CountyName = strings.TrimSpace(body.CountyName)
	body.CountyCode = strings.TrimSpace(body.CountyCode)

	return validation.ValidateStruct(body, validation.Field(&body.CountyName, validation.Required),
		validation.Field(&body.CountyCode, validation.Required, validation.Length(1, 3)))
}

func (rs Resource) HandleGetCounties(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleCreateCounty(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleUpdateCounty(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleDeleteCounty(w http.ResponseWriter, r *http.Request) {
}
