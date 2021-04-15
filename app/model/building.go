package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BuildingType string

const (
	Apartment   = BuildingType("apartment")
	Commercial  = BuildingType("commercial")
	Residential = BuildingType("residential")
	Onsale      = BuildingType("onsale")
)


type Building struct {
	BuildingId   uuid.UUID
	UserId       uuid.UUID
	BuildingName string       `gorm:"not null;unique"`
	BuildingType BuildingType `gorm:"column:building_type"`
	Description  string       `gorm:"not null"`
	Verified     bool         `gorm:"not null"`
	gorm.Model
}

type BuildingImage struct {
	BuildingImageId uuid.UUID
	BuildingId      uuid.UUID
	ImageUrl        string `gorm:"not null;unique"`
	gorm.Model
}
