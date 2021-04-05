package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	PaymentMethodId uuid.UUID `json:"payment_method_id"`
	Method          string    `json:"method"`
}
