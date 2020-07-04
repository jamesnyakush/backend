package model

type PaymentMethod struct {
	PaymentMethodId uint   `json:"payment_method_id"`
	Method          string `json:"method"`
}
