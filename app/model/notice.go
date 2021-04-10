package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Notice struct {
	NoticeId uuid.UUID
	gorm.Model
}
