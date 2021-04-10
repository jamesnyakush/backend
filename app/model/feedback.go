package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	FeedbackId  uuid.UUID
	Subject     string
	Description string
	gorm.Model
}
