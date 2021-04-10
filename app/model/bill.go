package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Bill struct {
	gorm.Model
	BillId                uuid.UUID
	BillingDate           time.Time
	LastBillingDate       time.Time
	Status                bool `gorm:"not null"`
	Range                 uint `gorm:"not null"`
	Description           uint `gorm:"not null"`
	Amount                uint `gorm:"not null"`
	UserId                uuid.UUID
	TransactionExternalId uuid.UUID
}
