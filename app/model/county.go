package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type County struct {
	gorm.Model
	CountyId   uuid.UUID `json:"county_id"`
	CountyName string    `json:"county_name"`
	CountyCode uint      `json:"county_code"`
}
