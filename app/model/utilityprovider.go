package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// We need this table so that we can store utility providers
// They must be enrolled for future purposes
type UtilityProvider struct {
	UtilityProviderId uuid.UUID
	Company           string `gorm:"not null;unique"`
	PhoneNumber       uint   `gorm:"not null;unique"`
	Email             string `gorm:"not null;unique"`
	gorm.Model
}

// some of the utility include [ Water, Waste , Wifi ] etc
type UtilityType struct {
	UtilityTypeId uuid.UUID
	Type          string `gorm:"not null;unique"`
	gorm.Model
}
