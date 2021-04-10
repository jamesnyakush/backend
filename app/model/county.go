package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type County struct {
	gorm.Model
	CountyId   uuid.UUID
	CountyName string `gorm:"not null;unique"`
	CountyCode uint   `gorm:"not null;unique"`
}
