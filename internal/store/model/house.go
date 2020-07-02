package model

import "time"

type House struct {
	HouseId        uint         `json:"house_id"`
	UserId         uint         `json:"user_id"`
	BuildingId     uint         `json:"building_id"`
	BuildingTypeId uint         `json:"building_type_id"`
	HouseImages    []HouseImage `json:"house_images"`
	HouseNumber    string       `json:"house_number"`
	Title          string       `json:"title"`
	Room           int          `json:"room"`
	Description    string       `json:"description"`
	Verified       bool         `json:"verified"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      time.Time    `json:"deleted_at"`
}

type HouseImage struct {
	HouseImageId uint      `json:"house_image_id"`
	HouseId      uint      `json:"house_id"`
	ImageUrl     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type HouseType struct {
	HouseTypeId uint      `json:"house_type_id"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
