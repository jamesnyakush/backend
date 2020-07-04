package model

import "time"

type PromoCode struct {
	PromoCodeId        string    `json:"promo_code_id"`
	Code               string    `json:"code"`
	Message            string    `json:"message"`
	Users              uint      `json:"users"`
	Status             bool      `json:"status"`
	Discount           string    `json:"discount"`
	DiscountType       string    `json:"discount_type"`
	MinimumOrderAmount string    `json:"minimum_order_amount"`
	MaxDiscountAmount  string    `json:"max_discount_amount"`
	RepeatUsage        string    `json:"repeat_usage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
}
