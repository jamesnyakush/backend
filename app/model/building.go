package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Building struct {
	BuildingId   uuid.UUID
	UserId       uuid.UUID
	BuildingName string `gorm:"not null;unique"`
	Verified     bool   `gorm:"not null"`
	Description  string `gorm:"not null"`
	gorm.Model
}

type BuildingImage struct {
	BuildingImageId uuid.UUID
	BuildingId      uuid.UUID
	ImageUrl        string `gorm:"not null;unique"`
	gorm.Model
}

type BuildingType struct {
	BuildingTypeId uuid.UUID
	Type           string `gorm:"not null; unique"`
	gorm.Model
}
