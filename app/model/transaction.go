package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	TransactionId     uuid.UUID
	BillId            uuid.UUID
	TransactionAmount uint
	TransactionDate   time.Time
	PaymentMethodId   uint
	gorm.Model
}
