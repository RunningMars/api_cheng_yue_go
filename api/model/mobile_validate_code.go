package model

import "time"

type MobileValidateCode struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Mobile    string    `json:"mobile"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
}

func (MobileValidateCode) TableName() string {
	return "mobile_validate_code"
}
