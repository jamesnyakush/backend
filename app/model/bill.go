package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Bill struct {
	BillId                uuid.UUID `json:"bill_id"`
	UserId                uint      `json:"user_id"`
	BillingDate           time.Time `json:"billing_date"`
	LastBillingDate       time.Time `json:"last_billing_date"`
	Status                bool      `json:"status"`
	Range                 uint      `json:"range"`
	Description           uint      `json:"description"`
	Amount                uint      `json:"amount"`
	TransactionExternalId uint      `json:"transaction_external_id"`
}
