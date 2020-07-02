package model

import "time"

type Bill struct {
	BillId          uint      `json:"bill_id"`
	UserId          uint      `json:"user_id"`
	BillingDate     time.Time `json:"billing_date"`
	LastBillingDate time.Time `json:"last_billing_date"`
	Range           uint      `json:"range"`
	Description     uint      `json:"description"`
	Amount          uint      `json:"amount"`
}
