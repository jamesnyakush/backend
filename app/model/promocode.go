package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type PromoCode struct {
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
	gorm.Model
}
