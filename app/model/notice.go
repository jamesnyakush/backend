package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Notice struct {
	NoticeId    uuid.UUID
	BuildingID  uuid.UUID
	Description string `gorm:"not null"`
	gorm.Model
}
