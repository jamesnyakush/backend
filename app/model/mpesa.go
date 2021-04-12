package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type MpesaC2B struct {
	MpesaC2BID        uuid.UUID
	TransactionType   string `gorm:"not null"`
	TransactionID     string `gorm:"not null;unique"`
	TransactionTime   string `gorm:"not null"`
	TransactionAmount string `gorm:"not null"`
	MSISDN            string `gorm:"not null"`
	ShortCode         string `gorm:"not null"`
	OrgAccountBalance string `gorm:"not null"`
	BillRefNumber     string `gorm:"not null;unique"`
	InvoiceNumber     string `gorm:"not null;unique"`
	FirstName         string `gorm:"not null"`
	MiddleName        string `gorm:"not null"`
	LastName          string `gorm:"not null"`
	gorm.Model
}

type MpesaB2C struct {
	MpesaB2CID                          uuid.UUID
	ResultCode                          string `gorm:"not null"`
	ResultDesc                          string `gorm:"not null"`
	OriginatorConversationID            string `gorm:"not null"`
	ConversationID                      string `gorm:"not null"`
	TransactionID                       string `gorm:"not null;unique"`
	Msisdn                              string `gorm:"not null"`
	B2CWorkingAccountAvailableFunds     string `gorm:"not null"`
	B2CUtilityAccountAvailableFunds     string `gorm:"not null"`
	Amount                              string `gorm:"not null"`
	TransactionCompletedDateTime        string `gorm:"not null"`
	B2CChargesPaidAccountAvailableFunds string `gorm:"not null"`
	ReceiverPartyPublicName             string `gorm:"not null"`
	B2CRecipientIsRegisteredCustomer    string `gorm:"not null"`
	gorm.Model
}
