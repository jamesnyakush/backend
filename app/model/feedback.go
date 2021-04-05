package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	FeedbackId  uuid.UUID `json:"feedback_id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
}
