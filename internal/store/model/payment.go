package model

import "github.com/gofrs/uuid"

type PaymentMethod struct {
	PaymentMethodId uuid.UUID `json:"payment_method_id"`
	Method          string    `json:"method"`
}
