package bill

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nyumbapoa/backend/app/errors"
)

type BillParams struct {
	BillingDate           string `json:"billing_date" schema:"billing_date" form:"billing_date"`
	LastBillingDate       string `json:"last_billing_date" schema:"last_billing_date" form:"last_billing_date"`
	Status                string `json:"status" schema:"status" form:"status"`
	Range                 string `json:"range" schema:"range" form:"range"`
	Description           string `json:"description" schema:"description" form:"description"`
	Amount                string `json:"amount" schema:"amount" form:"amount"`
	TransactionExternalId string `json:"transaction_external_id" schema:"transaction_external_id" form:"transaction_external_id"`
}

func (body BillParams) Validate() error {
	err := validation.ValidateStruct(body,
		validation.Field(&body.BillingDate, validation.Required, is.ASCII),
		validation.Field(&body.LastBillingDate, validation.Required, is.ASCII),
		validation.Field(&body.Status, validation.Required, is.Int),
		validation.Field(&body.Range, validation.Required, is.Int),
		validation.Field(&body.Description, validation.Length(8, 32), is.Alphanumeric),
		validation.Field(&body.Amount, validation.Required, is.Float),
		validation.Field(&body.TransactionExternalId, validation.Required, is.UUID),
	)

	return errors.ParseValidationErrorMap(err)
}

