package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type LandType string
type SoilType string

type Land struct {
	ID        uuid.UUID
	LandType  string
	TitleType string
	Price     string
	Size      int
	SoilType  string
	Location  string
	gorm.Model
}
