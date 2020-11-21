package bill

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"strings"
)

type billRequest struct {
	UserId                string `json:"user_id"`
	BillingDate           string `json:"billing_date"`
	LastBillingDate       string `json:"last_billing_date"`
	Status                string `json:"status"`
	Range                 string `json:"range"`
	Description           string `json:"description"`
	Amount                string `json:"amount"`
	TransactionExternalId string `json:"transaction_external_id"`
}

func (body *billRequest) Bind(r *http.Request) error {
	body.UserId = strings.TrimSpace(body.UserId)
	body.BillingDate = strings.TrimSpace(body.BillingDate)
	body.LastBillingDate = strings.TrimSpace(body.LastBillingDate)
	body.Status = strings.TrimSpace(body.Status)
	body.Range = strings.TrimSpace(body.Range)
	body.Description = strings.TrimSpace(body.Description)
	body.Amount = strings.TrimSpace(body.Amount)
	body.TransactionExternalId = strings.TrimSpace(body.TransactionExternalId)

	return validation.ValidateStruct(body, validation.Field(&body.UserId, validation.Required, is.Alphanumeric),
		validation.Field(&body.BillingDate, validation.Required, is.ASCII), validation.Field(&body.LastBillingDate, validation.Required, is.ASCII),
		validation.Field(&body.Status, validation.Required, is.Int), validation.Field(&body.Range, validation.Required, is.Int), validation.Field(&body.Description, validation.Length(8, 32), is.Alphanumeric),
		validation.Field(&body.Amount, validation.Required, is.Float), validation.Field(&body.TransactionExternalId, validation.Required, is.UUID))
}

func (rs Resource) HandleGetBill(w http.ResponseWriter, r *http.Request) {

	rs.Store.Bill().Create()
}

func (rs Resource) HandleGetBills(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleCreateBill(w http.ResponseWriter, r *http.Request) {
}

func (rs Resource) HandleUpdateBill(w http.ResponseWriter, r *http.Request) {
}
