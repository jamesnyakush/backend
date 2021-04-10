package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type House struct {
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
	gorm.Model
}

type HouseImage struct {
	HouseImageId uuid.UUID
	HouseId      uuid.UUID
	ImageUrl     string
	gorm.Model
}

type HouseType struct {
	HouseTypeId uuid.UUID
	Type        string `gorm:"not null;unique"`
	gorm.Model
}
