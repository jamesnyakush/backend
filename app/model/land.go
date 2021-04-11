package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Land struct {
	LandID    uuid.UUID
	LandType  string
	TitleType string
	Price     string
	Size      int
	SoilType  string
	Location  string
	gorm.Model
}
