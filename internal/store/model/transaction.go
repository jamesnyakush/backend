package model

import "time"

type Transaction struct {
	TransactionId uint `json:"transaction_id"`
	BillId uint `json:"bill_id"`
	TransactionAmount uint `json:"transaction_amount"`
	TransactionDate time.Time `json:"transaction_date"`
	PaymentMethodId uint `json:"payment_method_id"`
}
