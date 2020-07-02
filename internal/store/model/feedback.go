package model

type Feedback struct {
	FeedbackId  string `json:"feedback_id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}
