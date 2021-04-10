package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	FeedbackId  uuid.UUID
	Subject     string
	Description string
}
