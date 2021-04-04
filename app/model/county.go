package model

import "github.com/gofrs/uuid"

type County struct {
	CountyId   uuid.UUID `json:"county_id"`
	CountyName string    `json:"county_name"`
	CountyCode uint      `json:"county_code"`
}
