package promocode

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nyumbapoa/backend/app/errors"
)

type promoCodeRequest struct {
	Code               string `json:"code"`
	Message            string `json:"message"`
	Status             string `json:"status"`
	Discount           string `json:"discount"`
	DiscountType       string `json:"discount_type"`
	MinimumOrderAmount string `json:"minimum_order_amount"`
	MaxDiscountAmount  string `json:"max_discount_amount"`
	RepeatUsage        string `json:"repeat_usage"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
}

func (body *promoCodeRequest) Validate() error {

	err := validation.ValidateStruct(body, validation.Field(&body.Code, validation.Required, validation.Length(8, 16)),
		validation.Field(&body.Message, validation.Length(8, 16)), validation.Field(&body.Discount, validation.Required),
		validation.Field(&body.DiscountType, validation.Required), validation.Field(&body.MinimumOrderAmount, validation.Required),
		validation.Field(&body.MaxDiscountAmount, validation.Required), validation.Field(&body.RepeatUsage, validation.Required),
		validation.Field(&body.StartDate, validation.Required), validation.Field(&body.EndDate, validation.Required))

	return errors.ParseValidationErrorMap(err)
}
