package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type HouseType string

type House struct {
	ID          uuid.UUID
	Number      string    `gorm:"not null;unique"`
	Title       string    `gorm:"not null;unique"`
	Room        uint      `gorm:"not null"`
	HouseType   HouseType `gorm:"column:house_type"`
	Verified    bool      `gorm:"not null"`
	Occupied    bool      `gorm:"not null"`
	Description string    `gorm:"not null"`
	UserId      uuid.UUID
	BuildingId  uuid.UUID
	gorm.Model
}

type HouseImage struct {
	ID       uuid.UUID
	HouseId  uuid.UUID
	ImageUrl string
	gorm.Model
}
