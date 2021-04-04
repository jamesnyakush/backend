package model

import "github.com/gofrs/uuid"

type UtilityProvider struct {
	UtilityProviderId uuid.UUID `json:"utility_provider_id"`
	Company           string    `json:"company"`
	CountyCode        string    `json:"county_code"`
	CountyName        string    `json:"county_name"`
	PhoneNumber       uint      `json:"phone_number"`
	Email             string    `json:"email"`
}

type UtilityType struct {
	UtilityTypeId uuid.UUID `json:"utility_type_id"`
	Type          string    `json:"type"`
}
