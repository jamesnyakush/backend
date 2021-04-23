package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	ID          uuid.UUID
	Subject     string
	Description string
	gorm.Model
}
