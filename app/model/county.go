package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type County struct {
	ID         uuid.UUID
	CountyName string `gorm:"not null;unique"`
	CountyCode uint   `gorm:"not null;unique"`
	gorm.Model
}
