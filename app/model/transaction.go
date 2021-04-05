package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	TransactionId     uuid.UUID `json:"transaction_id"`
	BillId            uint      `json:"bill_id"`
	TransactionAmount uint      `json:"transaction_amount"`
	TransactionDate   time.Time `json:"transaction_date"`
	PaymentMethodId   uint      `json:"payment_method_id"`
}
