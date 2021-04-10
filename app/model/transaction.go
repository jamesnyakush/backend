package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	TransactionId     uuid.UUID
	BillId            uuid.UUID
	TransactionAmount uint
	TransactionDate   time.Time
	PaymentMethodId   uint
}
