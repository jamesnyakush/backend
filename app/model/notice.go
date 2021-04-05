package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	NoticeId uuid.UUID `json:"notice_id"`
}
