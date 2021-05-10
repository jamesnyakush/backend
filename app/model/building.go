package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Building struct {
	ID             uuid.UUID
	UserId         uuid.UUID
	BuildingName   string `gorm:"not null;unique"`
	BuildingTypeId uuid.UUID
	Description    string `gorm:"not null"`
	Verified       bool   `gorm:"not null"`
	gorm.Model
}

type BuildingType struct {
	ID   uuid.UUID
	Type string `gorm:"not null;unique"`
}
type BuildingImage struct {
	ID         uuid.UUID
	BuildingId uuid.UUID
	ImageUrl   string `gorm:"not null;unique"`
	gorm.Model
}
