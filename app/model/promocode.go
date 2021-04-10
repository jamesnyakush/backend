package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type PromoCode struct {
	gorm.Model
	PromoCodeId        uuid.UUID
	Code               string
	Message            string
	Status             bool
	Discount           string
	DiscountType       string
	MinimumOrderAmount string
	MaxDiscountAmount  string
	RepeatUsage        string
	StartDate          time.Time
	EndDate            time.Time
}
