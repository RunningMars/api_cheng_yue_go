package model

import "time"

type SendSMS struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	Mobile        string    `json:"mobile"`
	TemplateCode  string    `json:"template_code"`
	SignName      string    `json:"sign_name"`
	TemplateParam string    `json:"template_param"`
	Message       string    `json:"message"`
	Response      string    `json:"response"`
	Code          string    `json:"code"`
	BizID         string    `json:"biz_id"`
	RequestID     string    `json:"request_id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     int       `json:"created_by"`
}

func (SendSMS) TableName() string {
	return "send_sms"
}
