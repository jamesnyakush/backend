package model

import (
	"gorm.io/gorm"
	"time"
)

type Building struct {
	gorm.Model
	BuildingId     uint            `json:"building_id"`
	UserId         uint            `json:"user_id"`
	BuildingImages []BuildingImage `json:"building_images"`
	BuildingName   string          `json:"building_name"`
	Verified       bool            `json:"verified"`
	Description    string          `json:"description"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      time.Time       `json:"deleted_at"`
}

type BuildingImage struct {
	gorm.Model
	BuildingImageId uint      `json:"building_image_id"`
	BuildingId      uint      `json:"building_id"`
	ImageUrl        string    `json:"image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type BuildingType struct {
	gorm.Model
	BuildingTypeId uint      `json:"building_type_id"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
