package model

type County struct {
	CountyId   uint   `json:"county_id"`
	CountyName string `json:"county_name"`
	CountyCode uint   `json:"county_code"`
}