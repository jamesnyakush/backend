package model

import "github.com/gofrs/uuid"

type Feedback struct {
	FeedbackId  uuid.UUID `json:"feedback_id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
}
