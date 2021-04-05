package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type PromoCode struct {
	gorm.Model
	PromoCodeId        uuid.UUID `json:"promo_code_id"`
	Code               string    `json:"code"`
	Message            string    `json:"message"`
	Users              []User    `json:"users"`
	Status             bool      `json:"status"`
	Discount           string    `json:"discount"`
	DiscountType       string    `json:"discount_type"`
	MinimumOrderAmount string    `json:"minimum_order_amount"`
	MaxDiscountAmount  string    `json:"max_discount_amount"`
	RepeatUsage        string    `json:"repeat_usage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
}
