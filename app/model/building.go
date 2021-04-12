package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Building struct {
	BuildingId     uint            `json:"building_id"`
	UserId         uint            `json:"user_id"`
	BuildingName   string          `json:"building_name"`
	Verified       bool            `json:"verified"`
	Description    string          `json:"description"`
	gorm.Model
}

type BuildingImage struct {
	BuildingImageId uuid.UUID
	BuildingId      uuid.UUID
	ImageUrl        string
	gorm.Model
}

type BuildingType struct {
	BuildingTypeId uuid.UUID
	Type           string  `gorm:"not null; unique"`
	gorm.Model
}
