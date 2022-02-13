package models

import (
	"time"
)

type User struct {
	ID           *string   `json:"id" gorm:"primary_key"`
	PhoneNumber  string    `json:"phone_number"`
	OTP          string    `json:"otp"`
	CreatedAt    time.Time `json:"created_at" sql:"type:timestamp without time zone" `
	OTPCreatedAt time.Time `json:"otp_created_at" sql:"type:timestamp without time zone" `
}
