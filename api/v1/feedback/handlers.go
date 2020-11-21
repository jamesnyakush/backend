package feedback

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"strings"
)

type feedbackRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func (body *feedbackRequest) Bind(r *http.Request) error {
	body.Subject = strings.TrimSpace(body.Subject)
	body.Description = strings.TrimSpace(body.Description)

	return validation.ValidateStruct(body, validation.Field(&body.Subject, validation.Required), validation.Field(&body.Description, validation.Required))
}

func (rs Resource) handleCreateFeedback(w http.ResponseWriter, r *http.Request) {
	body := feedbackRequest{}
	fmt.Println(body)
}
