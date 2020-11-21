package role

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"strings"
)

type roleRequest struct {
	RoleName string `json:"role_name"`
}

func (body *roleRequest) Bind(r *http.Request) error {
	body.RoleName = strings.TrimSpace(body.RoleName)

	return validation.ValidateStruct(body, validation.Field(&body.RoleName, validation.Required, is.Alphanumeric))
}

func (rs Resource) HandleCreateRole(w http.ResponseWriter,r *http.Request)  {

}