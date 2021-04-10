package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	HouseId     uuid.UUID
	HouseNumber string `gorm:"not null;unique"`
	Title       string `gorm:"not null;unique"`
	Room        uint   `gorm:"not null"`
	Verified    bool   `gorm:"not null"`
	Occupied    bool   `gorm:"not null"`
	Description string `gorm:"not null"`
	UserId      uuid.UUID
	BuildingId  uuid.UUID
	HouseTypeId uuid.UUID
}

type HouseImage struct {
	gorm.Model
	HouseImageId uuid.UUID
	HouseId      uuid.UUID
	ImageUrl     string
}

type HouseType struct {
	gorm.Model
	HouseTypeId uuid.UUID
	Type        string `gorm:"not null;unique"`
}
