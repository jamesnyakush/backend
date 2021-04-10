package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Building struct {
	gorm.Model
	BuildingId     uint            `json:"building_id"`
	UserId         uint            `json:"user_id"`
	BuildingName   string          `json:"building_name"`
	Verified       bool            `json:"verified"`
	Description    string          `json:"description"`
}

type BuildingImage struct {
	gorm.Model
	BuildingImageId uuid.UUID
	BuildingId      uuid.UUID
	ImageUrl        string
}

type BuildingType struct {
	gorm.Model
	BuildingTypeId uuid.UUID
	Type           string  `gorm:"not null; unique"`
}
